# funcval
golang runtime funcval extract

[![Go1.14](https://github.com/visualfc/funcval/workflows/Go1.14/badge.svg)](https://github.com/visualfc/funcval/actions/workflows/go114.yml)
[![Go1.15](https://github.com/visualfc/funcval/workflows/Go1.15/badge.svg)](https://github.com/visualfc/funcval/actions/workflows/go115.yml)
[![Go1.16](https://github.com/visualfc/funcval/workflows/Go1.16/badge.svg)](https://github.com/visualfc/funcval/actions/workflows/go116.yml)
[![Go1.17](https://github.com/visualfc/funcval/workflows/Go1.17/badge.svg)](https://github.com/visualfc/funcval/actions/workflows/go117.yml)
[![Go1.18](https://github.com/visualfc/funcval/workflows/Go1.18/badge.svg)](https://github.com/visualfc/funcval/actions/workflows/go118.yml)

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
