#!/bin/sh
# Copyright 2016 The LUCI Authors. All rights reserved.
# Use of this source code is governed under the Apache License, Version 2.0
# that can be found in the LICENSE file.

cd "${0%/*}"
svcdec -output s1server_dec.golden -type S1Server,S2Server ../../../internal/svctool/testdata
