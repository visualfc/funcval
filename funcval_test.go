package funcval_test

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/visualfc/funcval"
)

func loadIndex(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func TestLoadIndex(t *testing.T) {
	fn := loadIndex(10)

	// check closure env
	fv, n := funcval.Get(fn)
	if n != 0 {
		t.Fatalf("must 0, %v", n)
	}
	type Closure struct {
		funcval.FuncVal
		i *int
	}
	c := (*Closure)(unsafe.Pointer(fv))
	if *c.i != 10 {
		t.Fatalf("error env %v", *c.i)
	}

	// check closure
	if r := fn(); r != 11 {
		t.Fatalf("error call %v", r)
	}
}

func TestFuncInt(t *testing.T) {
	var i int
	fn := func() int {
		return i
	}

	// check closure env
	fv, n := funcval.Get(fn)
	if n != 0 {
		t.Fatalf("must 0, %v", n)
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
	if r := fn(); r != i {
		t.Fatalf("error call %v", r)
	}
}

func TestFuncSlice(t *testing.T) {
	env := []interface{}{"hello", 100}
	fn := func() []interface{} {
		return env
	}

	// check closure env
	fv, n := funcval.Get(fn)
	if n != 0 {
		t.Fatalf("must 0, %v", n)
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
	r := fn()
	if fmt.Sprint(r) != "[hello world]" {
		t.Fatalf("error call %v", r)
	}
}
