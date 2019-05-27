package api

import (
	"AdBlockProxy/libs/config"
	"net/http"
)

func Hosts(w http.ResponseWriter, r *http.Request) {
	hosts := ""
	for k := range config.Config.Blacklist {
		hosts += "192.168.0.100 " + k + "\n"
	}
	w.Write([]byte(hosts))
}
