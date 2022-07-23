package pk2

import (
	"github.com/wzbwzt/studyGo/init/pk3"
)

func init() {
	println("init pk2")
}

func Get() {
	println("get pk2")
	pk3.Get()
}
