package main

import (
	"AdBlockProxy/libs/config"
	"AdBlockProxy/libs/proxy"
	"flag"
	_ "github.com/liuzheng/golog"
)

const (
	Name    = "AdBlockProxy"
	version = "0.1"
)

func main() {
	flag.Parse()
	config.LoadConfig()
	service := proxy.Server{Addr: "0.0.0.0:8080"}
	service.Start()
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	//<-signals
}
