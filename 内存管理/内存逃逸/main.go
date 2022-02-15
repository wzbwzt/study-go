package main

import (
	"unsafe"
)

type A struct {
	S *string
}

func (f *A) String() string {
	return *f.S
}

type ATrick struct {
	S unsafe.Pointer
}

func (f *ATrick) String() string {
	return *(*string)(f.S)
}

func NewA(s string) A {
	return A{S: &s}
}

func NewATrick(s string) ATrick {
	return ATrick{S: noescape(unsafe.Pointer(&s))}
}

// noescape hides a pointer from escape analysis.  noescape is
// the identity function but escape analysis doesn't think the
// output depends on the input.  noescape is inlined and currently
// compiles down to zero instructions.
// USE CAREFULLY!
//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

func main() {
	s := "hello"
	f1 := NewA(s)
	f2 := NewATrick(s)
	s1 := f1.String()
	s2 := f2.String()
	_ = s1 + s2
}
