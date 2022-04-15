package main

import (
	"io/ioutil"

	"github.com/gookit/goutil/dump"
)

func main() {
	getAllFile()

}

func getAllFile() (err error) {
	fs, err := ioutil.ReadDir("../")
	if err != nil {
		return
	}
	for _, f := range fs {
		f.Mode()
		if f.IsDir() {
			dump.P("dir:", f.Name())
		} else {
			dump.P("file:", f.Name())
		}
	}
	return
}
