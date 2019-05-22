package main

import (
	_ "github.com/liuzheng/golog"
	"AdBlockProxy/libs/proxy"
	"flag"
)

const (
	Name    = "AdBlockProxy"
	version = "0.1"
)

func main() {
	flag.Parse()
	service := proxy.Server{Addr: "0.0.0.0:8080"}
	service.Start()
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	//<-signals
}
