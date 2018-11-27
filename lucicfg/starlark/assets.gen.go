// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// AUTOGENERATED. DO NOT EDIT.

// Package starlark is generated by go.chromium.org/luci/tools/cmd/assets.
//
// It contains all [*.css *.html *.js *.star *.tmpl] files found in the package as byte arrays.
package starlark

// GetAsset returns an asset by its name. Returns nil if no such asset exists.
func GetAsset(name string) []byte {
	return []byte(files[name])
}

// GetAssetString is version of GetAsset that returns string instead of byte
// slice. Returns empty string if no such asset exists.
func GetAssetString(name string) string {
	return files[name]
}

// GetAssetSHA256 returns the asset checksum. Returns nil if no such asset
// exists.
func GetAssetSHA256(name string) []byte {
	data := fileSha256s[name]
	if data == nil {
		return nil
	}
	return append([]byte(nil), data...)
}

// Assets returns a map of all assets.
func Assets() map[string]string {
	cpy := make(map[string]string, len(files))
	for k, v := range files {
		cpy[k] = v
	}
	return cpy
}

var files = map[string]string{
	"stdlib/builtins.star": string([]byte{35, 32,
		67, 111, 112, 121, 114, 105, 103, 104, 116, 32, 50, 48, 49, 56,
		32, 84, 104, 101, 32, 76, 85, 67, 73, 32, 65, 117, 116, 104,
		111, 114, 115, 46, 10, 35, 10, 35, 32, 76, 105, 99, 101, 110,
		115, 101, 100, 32, 117, 110, 100, 101, 114, 32, 116, 104, 101, 32,
		65, 112, 97, 99, 104, 101, 32, 76, 105, 99, 101, 110, 115, 101,
		44, 32, 86, 101, 114, 115, 105, 111, 110, 32, 50, 46, 48, 32,
		40, 116, 104, 101, 32, 34, 76, 105, 99, 101, 110, 115, 101, 34,
		41, 59, 10, 35, 32, 121, 111, 117, 32, 109, 97, 121, 32, 110,
		111, 116, 32, 117, 115, 101, 32, 116, 104, 105, 115, 32, 102, 105,
		108, 101, 32, 101, 120, 99, 101, 112, 116, 32, 105, 110, 32, 99,
		111, 109, 112, 108, 105, 97, 110, 99, 101, 32, 119, 105, 116, 104,
		32, 116, 104, 101, 32, 76, 105, 99, 101, 110, 115, 101, 46, 10,
		35, 32, 89, 111, 117, 32, 109, 97, 121, 32, 111, 98, 116, 97,
		105, 110, 32, 97, 32, 99, 111, 112, 121, 32, 111, 102, 32, 116,
		104, 101, 32, 76, 105, 99, 101, 110, 115, 101, 32, 97, 116, 10,
		35, 10, 35, 32, 32, 32, 32, 32, 32, 104, 116, 116, 112, 58,
		47, 47, 119, 119, 119, 46, 97, 112, 97, 99, 104, 101, 46, 111,
		114, 103, 47, 108, 105, 99, 101, 110, 115, 101, 115, 47, 76, 73,
		67, 69, 78, 83, 69, 45, 50, 46, 48, 10, 35, 10, 35, 32,
		85, 110, 108, 101, 115, 115, 32, 114, 101, 113, 117, 105, 114, 101,
		100, 32, 98, 121, 32, 97, 112, 112, 108, 105, 99, 97, 98, 108,
		101, 32, 108, 97, 119, 32, 111, 114, 32, 97, 103, 114, 101, 101,
		100, 32, 116, 111, 32, 105, 110, 32, 119, 114, 105, 116, 105, 110,
		103, 44, 32, 115, 111, 102, 116, 119, 97, 114, 101, 10, 35, 32,
		100, 105, 115, 116, 114, 105, 98, 117, 116, 101, 100, 32, 117, 110,
		100, 101, 114, 32, 116, 104, 101, 32, 76, 105, 99, 101, 110, 115,
		101, 32, 105, 115, 32, 100, 105, 115, 116, 114, 105, 98, 117, 116,
		101, 100, 32, 111, 110, 32, 97, 110, 32, 34, 65, 83, 32, 73,
		83, 34, 32, 66, 65, 83, 73, 83, 44, 10, 35, 32, 87, 73,
		84, 72, 79, 85, 84, 32, 87, 65, 82, 82, 65, 78, 84, 73,
		69, 83, 32, 79, 82, 32, 67, 79, 78, 68, 73, 84, 73, 79,
		78, 83, 32, 79, 70, 32, 65, 78, 89, 32, 75, 73, 78, 68,
		44, 32, 101, 105, 116, 104, 101, 114, 32, 101, 120, 112, 114, 101,
		115, 115, 32, 111, 114, 32, 105, 109, 112, 108, 105, 101, 100, 46,
		10, 35, 32, 83, 101, 101, 32, 116, 104, 101, 32, 76, 105, 99,
		101, 110, 115, 101, 32, 102, 111, 114, 32, 116, 104, 101, 32, 115,
		112, 101, 99, 105, 102, 105, 99, 32, 108, 97, 110, 103, 117, 97,
		103, 101, 32, 103, 111, 118, 101, 114, 110, 105, 110, 103, 32, 112,
		101, 114, 109, 105, 115, 115, 105, 111, 110, 115, 32, 97, 110, 100,
		10, 35, 32, 108, 105, 109, 105, 116, 97, 116, 105, 111, 110, 115,
		32, 117, 110, 100, 101, 114, 32, 116, 104, 101, 32, 76, 105, 99,
		101, 110, 115, 101, 46, 10, 10, 100, 101, 102, 32, 115, 97, 121,
		95, 104, 105, 40, 109, 115, 103, 41, 58, 10, 32, 32, 95, 95,
		110, 97, 116, 105, 118, 101, 95, 95, 46, 101, 109, 105, 116, 95,
		103, 114, 101, 101, 116, 105, 110, 103, 40, 109, 115, 103, 41, 10,
	}),
}

var fileSha256s = map[string][]byte{
	"stdlib/builtins.star": {14, 67,
		69, 142, 231, 5, 47, 226, 233, 175, 144, 197, 0, 55, 159, 205,
		167, 144, 32, 38, 18, 52, 225, 95, 193, 19, 243, 159, 53, 106,
		92, 235},
}