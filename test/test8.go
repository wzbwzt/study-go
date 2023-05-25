//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	t := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-t.C:
			d := time.Since(start).Seconds()
			fmt.Printf("\b\b\b%03d", int(d))
		}
	}
}
