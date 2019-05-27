package main

import (
	"AdBlockProxy/libs/api"
	"AdBlockProxy/libs/config"
	"AdBlockProxy/libs/proxy"
	"flag"
	"github.com/liuzheng/golog"
	_ "github.com/liuzheng/golog"
	"net/http"
	"strconv"
)

const (
	Name    = "AdBlockProxy"
	version = "0.1"
)

func main() {
	flag.Parse()
	config.LoadConfig()

	addr := config.Config.Server + ":" + strconv.Itoa(int(config.Config.Listen))
	golog.Info(Name, "Start service at %v", addr)

	http.HandleFunc("/", proxy.ProxyHandler)
	http.HandleFunc("/adblockproxy/v1/loadconfig", api.LoadConfig)
	http.HandleFunc("/adblockproxy/v1/dumpconfig", api.DumpConfig)
	http.ListenAndServe(addr, nil)

	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	//<-signals
	//config.DumpConfig()
}
