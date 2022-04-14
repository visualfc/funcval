package funcval

import (
	"reflect"
	"unsafe"
)

type FuncVal struct {
	fn uintptr
	// variable-size, fn-specific data here
}

type eface struct {
	_type unsafe.Pointer
	word  unsafe.Pointer
}

var (
	dummy = reflect.MakeFunc(reflect.TypeOf((*func())(nil)).Elem(), nil).Pointer()
)

// Get is get func interface funcval and check make by reflect.MakeFunc
func Get(fn interface{}) (fv *FuncVal, makefunc bool) {
	v := (*eface)(unsafe.Pointer(&fn))
	if *(*uintptr)(v.word) == dummy {
		impl := (*makeFuncImpl)(v.word)
		return *(**FuncVal)(unsafe.Pointer(&impl.fn)), true
	}
	return (*FuncVal)(v.word), false
}

// reflect.makeFuncImpl
type makeFuncImpl struct {
	code   uintptr
	stack  unsafe.Pointer // ptrmap for both args and results
	argLen uintptr        // just args
	ftyp   *unsafe.Pointer
	fn     func([]reflect.Value) []reflect.Value
}
