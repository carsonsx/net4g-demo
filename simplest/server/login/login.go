package login

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/simplest/server/msg"
	"github.com/carsonsx/net4g-demo/simplest/server/tools"
	"reflect"
)

var Dispatcher = net4g.NewDispatcher("login", 1)

func init() {
	//Dispatcher.AddHandler(login)
	Dispatcher.AddHandler(jsonLogin, reflect.TypeOf(&msg.UserLogin{}))
	Dispatcher.OnConnectionClosed(func(agent net4g.NetAgent) {
		log4g.Info("remove session in redis")
		var offline msg.UserOffline
		offline.UseId = agent.Session().GetInt("userid")
		Dispatcher.BroadcastAll(&offline)
	})
}

func login(agent net4g.NetAgent) {
	log4g.Debug("g[%d] - %s", tools.GetGID(), agent.Msg().(string))
	agent.Write("login")
}

func jsonLogin(agent net4g.NetAgent) {
	log4g.Debug("g[%d] - %v", tools.GetGID(), agent.Msg())

	userLogin := agent.Msg().(*msg.UserLogin)

	var result msg.UserLoginReply
	if userLogin.Username == "carsonsx" && userLogin.Password == "123456" {
		result.Code = 0
		result.Msg = "login success"
	} else {
		result.Code = 1
		result.Msg = "wrong username or password"
	}

	agent.Session().Set("userid", 1)
	agent.Session().Set("username", userLogin.Username)

	log4g.Debug(result.Msg)

	agent.Write(&result)

	var online msg.UserOnline
	online.UseId = agent.Session().GetInt("userid")
	Dispatcher.BroadcastOthers(agent.Session(), &online)
}
