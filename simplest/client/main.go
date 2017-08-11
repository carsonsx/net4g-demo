package main

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"reflect"
	"github.com/carsonsx/net4g-demo/simplest/server/msg"
)

var serializer = net4g.NewJsonSerializer()

func main() {

	serializer.SerializeId(new(msg.UserLogin))
	serializer.DeserializeId(new(msg.UserLoginReply))
	serializer.DeserializeId(new(msg.UserOnline))
	serializer.DeserializeId(new(msg.UserOffline))
	serializer.SerializeId(new(msg.ChangeName))

	//log4g.SetLevel(log4g.LEVEL_TRACE)

	dispatcher := net4g.NewDispatcher("client", 1)

	dispatcher.AddHandler(func(agent net4g.NetAgent) {
		userOnline := agent.Msg().(*msg.UserOnline)
		log4g.Info("user[%d] is online", userOnline.UseId)
	}, reflect.TypeOf(&msg.UserOnline{}))

	dispatcher.AddHandler(func(agent net4g.NetAgent) {
		userLoginReply := agent.Msg().(*msg.UserLoginReply)
		if userLoginReply.Code == 0 {
			changeName := new(msg.ChangeName)
			changeName.NewName = "NewName"
			log4g.Debug("request change name")
			agent.Write(changeName)
		} else {
			log4g.Error(userLoginReply.Msg)
		}
	}, reflect.TypeOf(&msg.UserLoginReply{}))

	dispatcher.AddHandler(func(agent net4g.NetAgent) {
		userOffline := agent.Msg().(*msg.UserOffline)
		log4g.Info("user[%d] was offline", userOffline.UseId)
	}, reflect.TypeOf(&msg.UserOffline{}))

	dispatcher.OnConnectionCreated(func(agent net4g.NetAgent) {
		var userLogin msg.UserLogin
		userLogin.Username = "carsonsx"
		userLogin.Password = "123456"
		agent.Write(&userLogin)
	})

	net4g.NewTcpClient(net4g.NewNetKeyAddrFn("c1", ":9093")).
		SetSerializer(serializer).
		AddDispatchers(dispatcher).
		EnableHeartbeat().
		Connect().
		Wait()
}
