package funcval

import (
	"errors"
	"reflect"
	"unsafe"
)

type FuncVal struct {
	fn uintptr
	// variable-size, fn-specific data here
}

type reflectValue struct {
	typ  unsafe.Pointer
	ptr  unsafe.Pointer
	flag uintptr
}

func Get(fn interface{}) (*FuncVal, error) {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return nil, errors.New("fn must be a function")
	}
	r := (*reflectValue)(unsafe.Pointer(&v))
	return (*FuncVal)(unsafe.Pointer(r.ptr)), nil
}
