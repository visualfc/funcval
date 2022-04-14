//go:build go1.17
// +build go1.17

package funcval

import (
	"reflect"
	"unsafe"

	"github.com/visualfc/funcval/abi"
)

type makeFuncImpl struct {
	makeFuncCtxt
	ftyp unsafe.Pointer
	fn   func([]reflect.Value) []reflect.Value
}

type makeFuncCtxt struct {
	fn      uintptr
	stack   unsafe.Pointer // ptrmap for both stack args and results
	argLen  uintptr        // just args
	regPtrs abi.IntArgRegBitmap
}
