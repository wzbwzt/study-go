package model

import (
	"context"
	"time"

	"wireDemo/client/proto"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
)

func Goodbye(c *gin.Context) {
	type GoodbyeReq struct {
		Msg string `json:"msg"`
	}
	var (
		param GoodbyeReq
		resp  Response
	)
	defer c.JSON(200, &resp)
	err := c.ShouldBindJSON(&param)
	if err != nil {
		logger.Error(err)
		resp.Code = 400
		return
	}
	req := api.GoodbyeReq{Content: param.Msg}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	rsp, err := proto.Goodbye.Goodbye(ctx, &req)
	if err != nil {
		logger.Error(err)
		resp.Code = 400
		return
	}
	resp.Code = 200
	resp.Data = rsp.Content
	return
}
