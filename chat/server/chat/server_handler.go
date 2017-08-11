package chat

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/chat/global"
)

var ServerDispatcher = net4g.NewDispatcher("chat-server")

func init() {
	ServerDispatcher.AddHandler(sendMessage, global.SEND_MESSAGE_KEY)
	ServerDispatcher.AddHandler(setUserInfo, new(global.SetUserInfo))
}

func sendMessage(agent net4g.NetAgent) {
	log4g.Info(agent.Msg())
	RouterClientDispatcher.BroadcastOne(agent.RawPack(), nil)
}

func setUserInfo(agent net4g.NetAgent) {
	log4g.Info(agent.Msg())
	u := agent.Msg().(*global.SetUserInfo)
	agent.Session().Set("username", u.Username)
	var reply global.SetUserInfoReply
	reply.Success = true
	agent.Write(&reply)
}