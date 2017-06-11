package global

import (
	"github.com/carsonsx/net4g"
	"reflect"
)

const (
	SEND_MESSAGE_KEY = "SendMessage"
)


var ClientSerializer = net4g.NewJsonSerializer()
var ServerSerializer = net4g.NewJsonSerializer()
var RouterSerializer = net4g.NewJsonSerializer()

func init() {
	ClientSerializer.RegisterKey(SetUserInfoType, true, "SetUserInfo")
	ClientSerializer.RegisterKey(SetUserInfoReplyType, true, "SetUserInfoReply")
	ClientSerializer.RegisterKey(SendMessageType, true, SEND_MESSAGE_KEY)

	ServerSerializer.RegisterKey(SetUserInfoType, true, "SetUserInfo")
	ServerSerializer.RegisterKey(SetUserInfoReplyType, true, "SetUserInfoReply")
	ServerSerializer.RegisterKey(SendMessageType, false, SEND_MESSAGE_KEY)

	RouterSerializer.RegisterKey(SendMessageType, false, SEND_MESSAGE_KEY)
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

var SendMessageType = reflect.TypeOf(&SendMessage{})
var SetUserInfoType = reflect.TypeOf(&SetUserInfo{})
var SetUserInfoReplyType = reflect.TypeOf(&SetUserInfoReply{})