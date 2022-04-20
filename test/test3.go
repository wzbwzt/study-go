//go:build ignore
// +build ignore

package main

import (
	"github.com/gookit/goutil/dump"
)

func main() {
	a := "We are happy."
	dump.P(replaceSpace(a))

}

//双指针
func replaceSpace(s string) string {
	bs := []byte(s)
	length := len(bs)

	var spaceCount int
	for _, v := range bs {
		if v == ' ' {
			spaceCount++
		}
	}
	if spaceCount == 0 {
		return s
	}
	appendCount := spaceCount * 2
	tmp := make([]byte, appendCount)
	bs = append(bs, tmp...)
	i := length
	j := len(bs) - 1
	for i >= 0 {
		if bs[i] == ' ' {
			bs[j] = '0'
			bs[j-1] = '2'
			bs[j-2] = '%'
			i--
			j -= 2
		} else {
			bs[j] = bs[i]
			i--
			j--
		}
	}
	return string(bs)

}

func trip() {
}
