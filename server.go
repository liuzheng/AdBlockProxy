package main

import (
	"AdBlockProxy/libs/config"
	"AdBlockProxy/libs/proxy"
	"flag"
	_ "github.com/liuzheng/golog"
	"os"
	"os/signal"
	"syscall"
)

const (
	Name    = "AdBlockProxy"
	version = "0.1"
)

func main() {
	flag.Parse()
	config.LoadConfig()
	go proxy.Start()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	config.DumpConfig()
}
