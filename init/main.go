package main

import (
	"github.com/wzbwzt/studyGo/init/pk1"
	"github.com/wzbwzt/studyGo/init/pk2"
	_ "github.com/wzbwzt/studyGo/init/pk3"
)

func main() {
	pk1.Get()
	pk2.Get()
}
