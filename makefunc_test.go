package funcval_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/visualfc/funcval"
)

func TestMakeFuncInt(t *testing.T) {
	var i int
	fn := func([]reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf(i)}
	}
	vfn := reflect.MakeFunc(reflect.TypeOf((*func() int)(nil)).Elem(), fn)
	fv, err := funcval.Get(fn)
	if err != nil {
		t.Fatal(err)
	}
	type Closure struct {
		funcval.FuncVal
		i *int
	}
	i = 10
	c := (*Closure)(unsafe.Pointer(fv))
	if *c.i != 10 {
		t.Fatalf("error env %v", *c.i)
	}
	if r := vfn.Interface().(func() int)(); r != i {
		t.Fatalf("error call %v", r)
	}
}

func TestMakeFuncSlice(t *testing.T) {
	env := []interface{}{"hello", 100}
	fn := func([]reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf(env)}
	}
	vfn := reflect.MakeFunc(reflect.TypeOf((*func() []interface{})(nil)).Elem(), fn)
	fv, err := funcval.Get(fn)
	if err != nil {
		t.Fatal(err)
	}
	env[1] = "world"
	type Closure struct {
		funcval.FuncVal
		env []interface{}
	}
	c := (*Closure)(unsafe.Pointer(fv))
	if fmt.Sprint(c.env) != "[hello world]" {
		t.Fatalf("error env %v", c.env)
	}
	r := vfn.Interface().(func() []interface{})()
	if fmt.Sprint(r) != "[hello world]" {
		t.Fatalf("error call %v", r)
	}
}
