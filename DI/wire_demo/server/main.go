package main

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	err := config.Load(file.NewSource(file.WithPath("./../config.yaml")))
	if err != nil {
		logger.Error(err)
		return
	}
	app, err := initApp()
	if err != nil {
		logger.Error(err)
		return
	}
	err = app.Run()
	if err != nil {
		logger.Error(err)
		return
	}
}
