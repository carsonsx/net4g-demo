package msg

import (
	"github.com/carsonsx/net4g"
	"reflect"
)

var Serializer = net4g.NewJsonSerializer()

func init() {
	//jsonByKey()
	jsonById()
}

func jsonByKey() {
	Serializer.RegisterKey(reflect.TypeOf(&UserLogin{}), true)
	Serializer.RegisterKey(reflect.TypeOf(&UserLoginReply{}), true, "user_login_reply")
	Serializer.RegisterKey(reflect.TypeOf(&UserOnline{}), true)
	Serializer.RegisterKey(reflect.TypeOf(&UserOffline{}), true)
	Serializer.RegisterKey(reflect.TypeOf(&ChangeName{}), true)
}

func jsonById() {
	Serializer.RegisterId(reflect.TypeOf(&UserLogin{}), true)
	Serializer.RegisterId(reflect.TypeOf(&UserLoginReply{}), true)
	Serializer.RegisterId(reflect.TypeOf(&UserOnline{}), true)
	Serializer.RegisterId(reflect.TypeOf(&UserOffline{}), true)
	Serializer.RegisterId(reflect.TypeOf(&ChangeName{}), true)
}
