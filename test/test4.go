//go:build ignore
// +build ignore

package main

// import "C"

func main() {
	var ch chan struct{}
	<-ch
}
