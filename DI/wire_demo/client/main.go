package main

import (
	"wireDemo/client/handle"
	"wireDemo/client/proto"

	"github.com/micro/go-micro/config/source/file"
	"gorm.io/gorm/logger"
	"honnef.co/go/tools/config"
)

func main() {
	err := config.Load(file.NewSource(file.WithPath("./../config.yaml")))
	if err != nil {
		logger.Error(err)
		return
	}
	proto.NewClient()
	app, err := initApp()
	if err != nil {
		logger.Error(err)
		return
	}
	handle.WebRoute(app.Web)
	err = app.Run()
	if err != nil {
		logger.Error(err)
		return
	}
}
