package main

import (
	"errors"
	"fmt"
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/chat/global"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"time"
)

var dispatcher = net4g.NewDispatcher("chat-client")

func init() {
	dispatcher.AddHandler(setUserInfoReply, global.SetUserInfoType)
}

func addrFn() (addr string, err error) {
	cli, err := api.NewClient(&api.Config{
		//Address: "consul-dev:8500",
		Address: "192.168.56.201:8500",
	})
	if err != nil {
		return "", err
	}
	se, _, err := cli.Health().Service("chatserver", "", true, nil)
	if err != nil {
		return "", err
	}
	l := len(se)
	if l == 0 {
		text := "not found any chat router service"
		log4g.Error(text)
		return "", errors.New(text)
	}
	log4g.Debug("found %d services", l)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	selectIndex := r.Intn(l)
	log4g.Debug("service select index: %d", selectIndex)
	one := se[selectIndex].Service
	return fmt.Sprintf("%s:%d", one.Address, one.Port), nil
}

func setUserInfoReply(req net4g.NetReq, res net4g.NetRes) {

	if req.Msg().(*global.SetUserInfoReply).Success {
		var m global.SendMessage
		m.Text = "hello world!"
		for {
			time.Sleep(1 * time.Second)
			if res.Write(&m) != nil {
				break
			}
		}
	}

}

func main() {

	dispatcher.AddHandler(func(req net4g.NetReq, res net4g.NetRes) {
		log4g.Debug(req.Msg())
	})

	tcpCli := net4g.NewTcpClient(addrFn).
		SetSerializer(global.Serializer).
		AddDispatchers(dispatcher).
		OnConnected(func(conn net4g.NetConn, client *net4g.TCPClient) {
			var u global.SetUserInfo
			u.Username = "carson"
			client.Write(&u)
		}).
		Start()

	tcpCli.Wait()
}
