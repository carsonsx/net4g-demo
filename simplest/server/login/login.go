package login

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/simplest/server/msg"
	"github.com/carsonsx/net4g-demo/simplest/server/tools"
	"reflect"
)

var Dispatcher = net4g.NewDispatcher("login")

func init() {
	//Dispatcher.AddHandler(login)
	Dispatcher.AddHandler(jsonLogin, reflect.TypeOf(&msg.UserLogin{}))
	Dispatcher.OnConnectionClosed(func(session net4g.NetSession) {
		log4g.Info("remove session in redis")
		var offline msg.UserOffline
		offline.UseId = session.GetInt("userid")
		Dispatcher.BroadcastAll(&offline)
	})
}

func login(req net4g.NetReq, res net4g.NetRes) {
	log4g.Debug("g[%d] - %s", tools.GetGID(), req.Msg().(string))
	res.Write("login")
}

func jsonLogin(req net4g.NetReq, res net4g.NetRes) {
	log4g.Debug("g[%d] - %v", tools.GetGID(), req.Msg())

	userLogin := req.Msg().(*msg.UserLogin)

	var result msg.UserLoginReply
	if userLogin.Username == "carsonsx" && userLogin.Password == "123456" {
		result.Code = 0
		result.Msg = "login success"
	} else {
		result.Code = 1
		result.Msg = "wrong username or password"
	}

	req.Session().Set("userid", 1)
	req.Session().Set("username", userLogin.Username)

	log4g.Debug(result.Msg)

	res.Write(&result)
}
