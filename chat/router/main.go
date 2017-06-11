package main

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/chat/global"
)

var dispatcher = net4g.NewDispatcher("chat-router", 1)

func init() {
	dispatcher.AddHandler(forwardMessage, global.SEND_MESSAGE_KEY)
}

func forwardMessage(agent net4g.NetAgent) {
	log4g.Info(agent.RawPack())
	dispatcher.BroadcastAll(agent.RawPack())
}

func main() {
	net4g.NewTcpServer("chat-route", ":9000").SetSerializer(global.RouterSerializer).AddDispatchers(dispatcher).Start().Wait()
}
