package model

import (
	"context"
	"time"
	"wireDemo/api"
	"wireDemo/client/proto"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
)

func Greeting(c *gin.Context) {
	type GreetReq struct {
		Msg string `json:"msg"`
	}
	var (
		param GreetReq
		resp  Response
	)
	defer c.JSON(200, &resp)
	err := c.ShouldBindJSON(&param)
	if err != nil {
		logger.Error(err)
		resp.Code = 400
		return
	}
	req := api.GreetingReq{Content: param.Msg}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	rsp, err := proto.Greet.Greeting(ctx, &req)
	if err != nil {
		logger.Error(err)
		resp.Code = 400
		return
	}
	resp.Code = 200
	resp.Data = rsp.Content
	return
}
