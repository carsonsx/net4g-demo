package game

import (
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/simplest/server/msg"
	"github.com/carsonsx/net4g-demo/simplest/server/tools"
	"reflect"
)

var Dispatcher = net4g.NewDispatcher("game")

func init() {
	Dispatcher.AddHandler(changeName, reflect.TypeOf(&msg.ChangeName{}))
	Dispatcher.OnConnectionClosed(func(session net4g.NetSession) {
		log4g.Info("save session data to db")
	})
	Dispatcher.OnDestroy(func() {
		log4g.Info("save game data to db")
	})
}

func changeName(req net4g.NetReq, res net4g.NetRes) {
	log4g.Info("userid: %d", req.Session().GetInt("userid"))
	log4g.Info("g[%d] - %s", tools.GetGID(), req.Msg().(*msg.ChangeName).NewName)
	//res.Write("load")
}
