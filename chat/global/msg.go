package global

import (
	"github.com/carsonsx/net4g"
)

const (
	SEND_MESSAGE_KEY = "SendMessage"
)


var Serializer = net4g.NewJsonSerializer()

func init() {
	Serializer.RegisterId(new(SetUserInfo), "SetUserInfo")
	Serializer.RegisterId(new(SetUserInfoReply), "SetUserInfoReply")
	Serializer.RegisterId(new(SendMessage), SEND_MESSAGE_KEY)
}

type SendMessage struct {
	Text string
	To string
}

type SetUserInfo struct {
	UserId   int
	Username string
}

type SetUserInfoReply struct {
	Success bool
}
