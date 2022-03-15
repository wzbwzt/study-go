package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/logger"
)

type Options struct {
	Svr micro.Service
	Web *gin.Engine
}

func (o *Options) Run() error {
	var err error
	err = o.Web.Run(config.Get("http", "addr").String("0.0.0.0:8000"))
	if err != nil {
		logger.Error(err)
		return err
	}
	err = o.Svr.Run()
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
