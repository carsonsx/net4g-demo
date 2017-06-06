package main

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"reflect"
	"github.com/carsonsx/net4g-demo/simplest/server/msg"
)

func main() {

	dispatcher := net4g.NewDispatcher("client")
	dispatcher.AddHandler(func(req net4g.NetReq, res net4g.NetRes) {
		log4g.Debug(req.Msg())
	})

	dispatcher.AddHandler(func(req net4g.NetReq, res net4g.NetRes) {
		userLoginReply := req.Msg().(*msg.UserLoginReply)
		if userLoginReply.Code == 0 {
			changeName := new(msg.ChangeName)
			changeName.NewName = "NewName"
			res.Write(changeName)
		} else {
			log4g.Error(userLoginReply.Msg)
		}
	}, reflect.TypeOf(&msg.UserLoginReply{}))

	dispatcher.AddHandler(func(req net4g.NetReq, res net4g.NetRes) {
		userOffline := req.Msg().(*msg.UserOffline)
		log4g.Info("user[%s] was offline", userOffline.UseId)
	}, reflect.TypeOf(&msg.UserOffline{}))

	jsonClient := net4g.NewTcpClient(net4g.AddrFn(":9093")).AddDispatchers(dispatcher).EnableHeartbeat().Start()
	var userLogin msg.UserLogin
	userLogin.Username = "carsonsx"
	userLogin.Password = "123456"
	jsonClient.Write(&userLogin)
	jsonClient.Wait()
}
