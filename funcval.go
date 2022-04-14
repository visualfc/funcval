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

var (
	dummy = reflect.MakeFunc(reflect.TypeOf((*func())(nil)).Elem(), nil).Pointer()
)

func Get(fn interface{}) (*FuncVal, error) {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return nil, errors.New("fn must be a function")
	}
	r := (*reflectValue)(unsafe.Pointer(&v))
	if v.Pointer() == dummy {
		return getUser(r.ptr)
	}
	return (*FuncVal)(r.ptr), nil
}

func getUser(ptr unsafe.Pointer) (*FuncVal, error) {
	return nil, errors.New("not impl")
}
