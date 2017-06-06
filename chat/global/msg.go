package global

import (
	"github.com/carsonsx/net4g"
	"reflect"
)

var Serializer = net4g.NewJsonSerializer()

func init() {
	Serializer.RegisterByKey(SendMessageType, "SendMessage")
	Serializer.RegisterByKey(SetUserInfoType, "SetUserInfo")
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