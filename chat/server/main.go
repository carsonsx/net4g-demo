package main

import (
	"errors"
	"fmt"
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/chat/global"
	"github.com/carsonsx/net4g-demo/chat/server/chat"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"time"
)

func addrFn() (addr string, err error) {
	cli, err := api.NewClient(&api.Config{
		Address: "consul-dev:8500",
		//Address: "192.168.56.201:8500",
	})
	if err != nil {
		log4g.Error(err)
		return "", err
	}
	se, _, err := cli.Health().Service("chatrouter", "", true, nil)
	if err != nil {
		log4g.Error(err)
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

func main() {
	chat.RouterClient = net4g.NewTcpClient(addrFn).SetSerializer(global.Serializer).AddDispatchers(chat.RouterClientDispatcher).Start()
	net4g.NewTcpServer("chat-server", ":8000").SetSerializer(global.Serializer).AddDispatchers(chat.ServerDispatcher).Start().Wait(chat.RouterClient)
}