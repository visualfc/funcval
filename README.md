# funcval
golang runtime funcval extract


### extract closure env variable

_func Get(fn interface{}) (fv *FuncVal, makefunc int)_

```
package main

import (
	"unsafe"

	"github.com/visualfc/funcval"
)

func main() {
	a := 1
	b := "hello"
	c := 10
	fn := func() {
		println(a, b, c)
	}
	c = 20
	fv, _ := funcval.Get(fn)
	type closure struct {
		funcval.FuncVal
		a int
		b string
		c *int
	}
	env := (*closure)(unsafe.Pointer(fv))
	println(env.a, env.b, *env.c)
}

```
