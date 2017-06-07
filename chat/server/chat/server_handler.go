package chat

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/chat/global"
)

var ServerDispatcher = net4g.NewDispatcher("chat-server")

func init() {
	ServerDispatcher.AddHandler(sendMessage, global.SendMessageType)
	ServerDispatcher.AddHandler(setUserInfo, global.SetUserInfoType)
}

func sendMessage(req net4g.NetReq, res net4g.NetRes) {
	log4g.Info(req.Msg())
	RouterClient.Write(req.Msg())
}

func setUserInfo(req net4g.NetReq, res net4g.NetRes) {
	u := req.Msg().(*global.SetUserInfo)
	req.Session().Set("username", u.Username)
	var reply global.SetUserInfoReply
	reply.Success = true
	res.Write(&reply)
}
