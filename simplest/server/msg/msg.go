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
	Serializer.RegisterByKey(reflect.TypeOf(&UserLogin{}))
	Serializer.RegisterByKey(reflect.TypeOf(&UserLoginReply{}), "user_login_reply")
	Serializer.RegisterByKey(reflect.TypeOf(&UserOffline{}))
	Serializer.RegisterByKey(reflect.TypeOf(&ChangeName{}))
}

func jsonById() {
	Serializer.RegisterById(reflect.TypeOf(&UserLogin{}))
	Serializer.RegisterById(reflect.TypeOf(&UserLoginReply{}))
	Serializer.RegisterById(reflect.TypeOf(&UserOffline{}))
	Serializer.RegisterById(reflect.TypeOf(&ChangeName{}))
}
