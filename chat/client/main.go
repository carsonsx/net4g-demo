package main

import (
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/chat/global"
	"time"
	"github.com/carsonsx/log4g"
)

var dispatcher = net4g.NewDispatcher("chat-client", 1)

func init() {
	dispatcher.AddHandler(setUserInfoReply, global.SetUserInfoReplyType)
}

//func addrFn() (addr string, err error) {
//	cli, err := api.NewClient(&api.Config{
//		//Address: "consul-dev:8500",
//		Address: "192.168.56.201:8500",
//	})
//	if err != nil {
//		return "", err
//	}
//	se, _, err := cli.Health().Service("chatserver", "", true, nil)
//	if err != nil {
//		return "", err
//	}
//	l := len(se)
//	if l == 0 {
//		text := "not found any chat router service"
//		log4g.Error(text)
//		return "", errors.New(text)
//	}
//	log4g.Debug("found %d services", l)
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	selectIndex := r.Intn(l)
//	log4g.Debug("service select index: %d", selectIndex)
//	one := se[selectIndex].Service
//	return fmt.Sprintf("%s:%d", one.Address, one.Port), nil
//}

func setUserInfoReply(agent net4g.NetAgent) {

	if agent.Msg().(*global.SetUserInfoReply).Success {
		var m global.SendMessage
		m.Text = "hello world!"
		var stop bool
		for !stop {
			time.Sleep(1 * time.Second)
			if err := dispatcher.One(&m, func(err error) {
				stop = true
			}); err != nil {
				break
			}
		}

	}

}

func main() {

	log4g.SetLevel(log4g.LEVEL_TRACE)

	dispatcher.OnConnectionCreated(func(agent net4g.NetAgent) {
		var u global.SetUserInfo
		u.Username = "carson4"
		dispatcher.BroadcastAll(&u)
	})

	net4g.NewTcpClient(net4g.NewNetAddrFn(":8000")).
		SetSerializer(global.ClientSerializer).
		AddDispatchers(dispatcher).
		Start().
		Wait()
}
