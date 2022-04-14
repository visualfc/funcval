package funcval_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/visualfc/funcval"
)

func TestFuncInt(t *testing.T) {
	var i int
	fn := func() int {
		return i
	}
	fv, b := funcval.Get(fn)
	if b != false {
		t.Fatal(b)
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
	if r := fn(); r != i {
		t.Fatalf("error call %v", r)
	}
}

func TestFuncSlice(t *testing.T) {
	env := []interface{}{"hello", 100}
	fn := func() []interface{} {
		return env
	}
	fv, b := funcval.Get(fn)
	if b != false {
		t.Fatal(b)
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
	r := fn()
	if fmt.Sprint(r) != "[hello world]" {
		t.Fatalf("error call %v", r)
	}
}
