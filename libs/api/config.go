package api

import (
	"AdBlockProxy/libs/config"
	"github.com/liuzheng/golog"
	"net/http"
)

const lname = "api"

func LoadConfig(w http.ResponseWriter, r *http.Request) {
	golog.Debug(lname, "LoadConfig")
	config.LoadConfig()
	w.Write([]byte("success"))
}
func DumpConfig(w http.ResponseWriter, r *http.Request) {
	golog.Debug(lname, "DumpConfig")
	config.DumpConfig()
	w.Write([]byte("success"))
}