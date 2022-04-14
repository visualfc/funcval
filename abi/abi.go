// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

// IntArgRegBitmap is a bitmap large enough to hold one bit per
// integer argument/return register.
type IntArgRegBitmap [(IntArgRegs + 7) / 8]uint8
