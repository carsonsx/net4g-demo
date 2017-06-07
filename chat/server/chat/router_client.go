package chat

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/chat/global"
)

var RouterClient *net4g.TCPClient
var RouterClientDispatcher = net4g.NewDispatcher("router-client")

func init() {
	RouterClientDispatcher.AddHandler(forward, global.SendMessageType)
}

func forward(req net4g.NetReq, res net4g.NetRes) {
	log4g.Info(req.Msg())
	ServerDispatcher.BroadcastAll(req.Msg())
}
