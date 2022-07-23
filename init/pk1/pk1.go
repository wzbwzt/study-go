package pk1

import "github.com/wzbwzt/studyGo/init/pk3"

func init() {
	println("init pk1")
}

func Get() {
	println("get pk1")
	pk3.Get()
}
