//go:build !llgo
// +build !llgo

package funcval

import "runtime"

const (
	IsSupport = runtime.Compiler == "gc"
)
