package funcval

import (
	"reflect"
	"unsafe"
)

type FuncVal struct {
	Fn uintptr
	// variable-size, fn-specific data here
}

type eface struct {
	_type unsafe.Pointer
	word  unsafe.Pointer
}

var (
	dummy = reflect.MakeFunc(reflect.TypeOf((*func())(nil)).Elem(), nil).Pointer()
)

// Get returns function/closure interface *FuncVal and count of make by reflect.MakeFunc
func Get(fn interface{}) (fv *FuncVal, makefunc int) {
	v := (*eface)(unsafe.Pointer(&fn))
	fv = (*FuncVal)(v.word)
	for fv.Fn == dummy {
		impl := (*makeFuncImpl)(unsafe.Pointer(fv))
		fv = *(**FuncVal)(unsafe.Pointer(&impl.fn))
		makefunc++
	}
	return
}
