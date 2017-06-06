package main

import (
	"fmt"
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/net4g-demo/simplest/server/game"
	"github.com/carsonsx/net4g-demo/simplest/server/login"
	"net/http"
	"runtime"
	"github.com/carsonsx/net4g-demo/simplest/server/msg"
)

func main() {

	if log4g.IsTraceEnabled() {
		log4g.Trace("[Entrance]total goroutine: %d", runtime.NumGoroutine())
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	log4g.Info("HTTP server listening on %d\n", 8080)
	go func() {
		log4g.Fatal(http.ListenAndServe(":8080", nil))
	}()

	if log4g.IsTraceEnabled() {
		log4g.Trace("[NotStarted]total goroutine: %d", runtime.NumGoroutine())
	}

	s1 := net4g.NewTcpServer("s1", ":9091").SetSerializer(msg.Serializer).AddDispatchers(net4g.NewDispatcher("s1")).Start()

	s2 := net4g.NewTcpServer("s2", ":9092").SetSerializer(msg.Serializer).AddDispatchers(net4g.NewDispatcher("s2")).Start()

	s3 := net4g.NewTcpServer("s3", ":9093").SetSerializer(msg.Serializer).AddDispatchers(login.Dispatcher, game.Dispatcher).EnableHeartbeat().Start()

	s4 := net4g.NewTcpServer("4", ":9094").SetSerializer(msg.Serializer).AddDispatchers(net4g.NewDispatcher("s4")).Start()

	s4.Wait(s1, s2, s3)

	if log4g.IsTraceEnabled() {
		log4g.Trace("[Closed]total goroutine: %d", runtime.NumGoroutine())
	}

}
