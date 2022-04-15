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

	// check closure env
	fv, n := funcval.Get(vfn.Interface())
	if n != 1 {
		t.Fatalf("must 1, %v", n)
	}
	i = 10
	type Closure struct {
		funcval.FuncVal
		i *int
	}
	c := (*Closure)(unsafe.Pointer(fv))
	if *c.i != 10 {
		t.Fatalf("error env %v", *c.i)
	}

	// check closure
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

	// check closure env
	fv, n := funcval.Get(vfn.Interface())
	if n != 1 {
		t.Fatalf("must 1, %v", n)
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

	// check closure
	r := vfn.Interface().(func() []interface{})()
	if fmt.Sprint(r) != "[hello world]" {
		t.Fatalf("error call %v", r)
	}
}

func TestMakeFuncByMakeFunc(t *testing.T) {
	env := []interface{}{"hello", 100}
	fn := reflect.MakeFunc(reflect.TypeOf((*func([]reflect.Value) []reflect.Value)(nil)).Elem(),
		func(args []reflect.Value) []reflect.Value {
			return []reflect.Value{reflect.ValueOf([]reflect.Value{reflect.ValueOf(env)})}
		})
	vfn := reflect.MakeFunc(reflect.TypeOf((*func() []interface{})(nil)).Elem(),
		fn.Interface().(func([]reflect.Value) []reflect.Value))

	// check closure env
	fv, n := funcval.Get(vfn.Interface())
	if n != 2 {
		t.Fatalf("must 2, %v", n)
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

	// check closure
	r := vfn.Interface().(func() []interface{})()
	if fmt.Sprint(r) != "[hello world]" {
		t.Fatalf("error call %v", r)
	}
}
