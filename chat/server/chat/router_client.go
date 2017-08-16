package chat

import (
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/chat/global"
)

var RouterClientDispatcher = net4g.NewDispatcher("router-client")

func init() {
	RouterClientDispatcher.AddHandler(forward, global.SEND_MESSAGE_KEY)
}

func forward(agent net4g.NetAgent) {
	ServerDispatcher.BroadcastAll(agent.RawPack())
}
