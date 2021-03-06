// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.regabireflect && (ppc64 || ppc64le) && go1.18
// +build goexperiment.regabireflect
// +build ppc64 ppc64le
// +build go1.18

package abi

const (
	// See abi_generic.go.

	// R3 - R10, R14 - R17.
	IntArgRegs = 12

	// F1 - F12.
	FloatArgRegs = 12

	EffectiveFloatRegSize = 8
)
