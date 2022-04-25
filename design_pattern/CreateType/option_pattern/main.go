package main

import "fmt"

//选项者模式：传不定参数

type HttpClient struct {
	TimeOut   int
	Maxidle   int
	ErrorFunc func()
}

type OptionFunc func(*HttpClient)
type OptionFuncs []OptionFunc

func WithTimeOut(time int) OptionFunc {
	return func(hc *HttpClient) {
		hc.TimeOut = time
	}
}

func WithMaxidle(idle int) OptionFunc {
	return func(hc *HttpClient) {
		hc.Maxidle = idle
	}
}

func WithErrorFunc(errfunc func()) OptionFunc {
	return func(hc *HttpClient) {
		hc.ErrorFunc = errfunc
	}
}

func (this OptionFuncs) apply(c *HttpClient) {
	for _, f := range this {
		f(c)
	}
}

func NewHttpClient(ops ...OptionFunc) *HttpClient {
	c := &HttpClient{}
	OptionFuncs(ops).apply(c)
	return c
}

func main() {
	c := NewHttpClient(WithTimeOut(20), WithMaxidle(12), WithErrorFunc(func() {
		panic("err")
	}))
	fmt.Printf("%v", c)
}
