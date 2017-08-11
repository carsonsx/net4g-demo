package game

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/simplest/server/msg"
	"github.com/carsonsx/net4g-demo/simplest/server/tools"
	"reflect"
)

var Dispatcher = net4g.NewDispatcher("game", 1)

func init() {
	Dispatcher.AddHandler(changeName, reflect.TypeOf(&msg.ChangeName{}))
	Dispatcher.OnConnectionClosed(func(agent net4g.NetAgent) {
		log4g.Info("save session data to db")
	})
	Dispatcher.OnDestroy(func() {
		log4g.Info("save game data to db")
	})
}

func changeName(agent net4g.NetAgent) {
	log4g.Info("userid: %d", agent.Session().GetInt("userid"))
	log4g.Info("g[%d] - %s", tools.GetGID(), agent.Msg().(*msg.ChangeName).NewName)
	//res.Write("load")
}
