package proto

import (
	"wireDemo/api"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config"
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
