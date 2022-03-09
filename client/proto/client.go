package proto

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config"
	"github.com/wu-xie-888/micro-wire/api"
)

var (
	ServerName = config.Get("name", "server").String("micro-wire-server")
	Greet      api.GreetService
	Goodbye    api.GoodbyeService
)

func NewClient() {
	Greet = api.NewGreetService(ServerName, client.DefaultClient)
	Goodbye = api.NewGoodbyeService(ServerName, client.DefaultClient)
}
