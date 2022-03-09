package model

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"github.com/wu-xie-888/micro-wire/api"
	"github.com/wu-xie-888/micro-wire/client/proto"
	"time"
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
