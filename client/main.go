package main

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/logger"
	"github.com/wu-xie-888/micro-wire/client/handle"
	"github.com/wu-xie-888/micro-wire/client/proto"
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
