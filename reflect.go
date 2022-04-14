//go:build !go1.17
// +build !go1.17

package funcval

import (
	"reflect"
	"unsafe"
)

// reflect.makeFuncImpl
type makeFuncImpl struct {
	code   uintptr
	stack  unsafe.Pointer // ptrmap for both args and results
	argLen uintptr        // just args
	ftyp   *unsafe.Pointer
	fn     func([]reflect.Value) []reflect.Value
}
