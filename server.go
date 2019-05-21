package main

import (
	"AdBlockProxy/libs/proxy"
	"flag"
)

const (
	Name    = "AdBlockProxy"
	version = "0.1"
)

func main() {
	flag.Parse()
	service := proxy.Server{Addr: "0.0.0.0:3000"}
	service.Start()
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	//<-signals
}
