// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

import (
	"bytes"
	"math"
	"strconv"
)

// appendFloat formats given float in bitSize, and appends to the given []byte.
func appendFloat(out []byte, n float64, bitSize int) []byte {
	switch {
	case math.IsNaN(n):
		return append(out, `"NaN"`...)
	case math.IsInf(n, +1):
		return append(out, `"Infinity"`...)
	case math.IsInf(n, -1):
		return append(out, `"-Infinity"`...)
	}

	// JSON number formatting logic based on encoding/json.
	// See floatEncoder.encode for reference.
	fmt := byte('f')
	if abs := math.Abs(n); abs != 0 {
		if bitSize == 64 && (abs < 1e-6 || abs >= 1e21) ||
			bitSize == 32 && (float32(abs) < 1e-6 || float32(abs) >= 1e21) {
			fmt = 'e'
		}
	}
	out = strconv.AppendFloat(out, n, fmt, -1, bitSize)
	if fmt == 'e' {
		n := len(out)
		if n >= 4 && out[n-4] == 'e' && out[n-3] == '-' && out[n-2] == '0' {
			out[n-2] = out[n-1]
			out = out[:n-1]
		}
	}
	return out
}

// consumeNumber reads the given []byte for a valid JSON number. If it is valid,
// it returns the number of bytes.  Parsing logic follows the definition in
// https://tools.ietf.org/html/rfc7159#section-6, and is based off
// encoding/json.isValidNumber function.
func consumeNumber(input []byte) (int, bool) {
	var n int

	s := input
	if len(s) == 0 {
		return 0, false
	}

	// Optional -
	if s[0] == '-' {
		s = s[1:]
		n++
		if len(s) == 0 {
			return 0, false
		}
	}

	// Digits
	switch {
	case s[0] == '0':
		s = s[1:]
		n++

	case '1' <= s[0] && s[0] <= '9':
		s = s[1:]
		n++
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
			n++
		}

	default:
		return 0, false
	}

	// . followed by 1 or more digits.
	if len(s) >= 2 && s[0] == '.' && '0' <= s[1] && s[1] <= '9' {
		s = s[2:]
		n += 2
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
			n++
		}
	}

	// e or E followed by an optional - or + and
	// 1 or more digits.
	if len(s) >= 2 && (s[0] == 'e' || s[0] == 'E') {
		s = s[1:]
		n++
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
			n++
			if len(s) == 0 {
				return 0, false
			}
		}
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
			n++
		}
	}

	// Check that next byte is a delimiter or it is at the end.
	if n < len(input) && isNotDelim(input[n]) {
		return 0, false
	}

	return n, true
}

// numberParts is the result of parsing out a valid JSON number. It contains
// the parts of a number. The parts are used for integer conversion.
type numberParts struct {
	neg  bool
	intp []byte
	frac []byte
	exp  []byte
}

// parseNumber constructs numberParts from given []byte. The logic here is
// similar to consumeNumber above with the difference of having to construct
// numberParts. The slice fields in numberParts are subslices of the input.
func parseNumber(input []byte) (numberParts, bool) {
	var neg bool
	var intp []byte
	var frac []byte
	var exp []byte

	s := input
	if len(s) == 0 {
		return numberParts{}, false
	}

	// Optional -
	if s[0] == '-' {
		neg = true
		s = s[1:]
		if len(s) == 0 {
			return numberParts{}, false
		}
	}

	// Digits
	switch {
	case s[0] == '0':
		// Skip first 0 and no need to store.
		s = s[1:]

	case '1' <= s[0] && s[0] <= '9':
		intp = s
		n := 1
		s = s[1:]
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
			n++
		}
		intp = intp[:n]

	default:
		return numberParts{}, false
	}

	// . followed by 1 or more digits.
	if len(s) >= 2 && s[0] == '.' && '0' <= s[1] && s[1] <= '9' {
		frac = s[1:]
		n := 1
		s = s[2:]
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
			n++
		}
		frac = frac[:n]
	}

	// e or E followed by an optional - or + and
	// 1 or more digits.
	if len(s) >= 2 && (s[0] == 'e' || s[0] == 'E') {
		s = s[1:]
		exp = s
		n := 0
		if s[0] == '+' || s[0] == '-' {
			s = s[1:]
			n++
			if len(s) == 0 {
				return numberParts{}, false
			}
		}
		for len(s) > 0 && '0' <= s[0] && s[0] <= '9' {
			s = s[1:]
			n++
		}
		exp = exp[:n]
	}

	return numberParts{
		neg:  neg,
		intp: intp,
		frac: bytes.TrimRight(frac, "0"), // Remove unnecessary 0s to the right.
		exp:  exp,
	}, true
}

// normalizeToIntString returns an integer string in normal form without the
// E-notation for given numberParts. It will return false if it is not an
// integer or if the exponent exceeds than max/min int value.
func normalizeToIntString(n numberParts) (string, bool) {
	intpSize := len(n.intp)
	fracSize := len(n.frac)

	if intpSize == 0 && fracSize == 0 {
		return "0", true
	}

	var exp int
	if len(n.exp) > 0 {
		i, err := strconv.ParseInt(string(n.exp), 10, 32)
		if err != nil {
			return "", false
		}
		exp = int(i)
	}

	var num []byte
	if exp >= 0 {
		// For positive E, shift fraction digits into integer part and also pad
		// with zeroes as needed.

		// If there are more digits in fraction than the E value, then the
		// number is not an integer.
		if fracSize > exp {
			return "", false
		}

		// Set cap to make a copy of integer part when appended.
		num = n.intp[:len(n.intp):len(n.intp)]
		num = append(num, n.frac...)
		for i := 0; i < exp-fracSize; i++ {
			num = append(num, '0')
		}

	} else {
		// For negative E, shift digits in integer part out.

		// If there are fractions, then the number is not an integer.
		if fracSize > 0 {
			return "", false
		}

		// index is where the decimal point will be after adjusting for negative
		// exponent.
		index := intpSize + exp
		if index < 0 {
			return "", false
		}

		num = n.intp
		// If any of the digits being shifted to the right of the decimal point
		// is non-zero, then the number is not an integer.
		for i := index; i < intpSize; i++ {
			if num[i] != '0' {
				return "", false
			}
		}
		num = num[:index]
	}

	if n.neg {
		return "-" + string(num), true
	}
	return string(num), true
}
