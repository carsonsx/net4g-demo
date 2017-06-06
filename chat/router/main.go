package main

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/chat/global"
)

var dispatcher = net4g.NewDispatcher("chat-router")

func init() {
	dispatcher.AddHandler(forwardMessage, global.SendMessageType)
}

func forwardMessage(req net4g.NetReq, res net4g.NetRes) {
	log4g.Info(req.Msg())
	dispatcher.BroadcastAll(req.Msg())
}

type Message struct {
	Text string
}

func main() {
	net4g.NewTcpServer("chat-route", ":9000").SetSerializer(global.Serializer).AddDispatchers(dispatcher).Start().Wait()
}
