package proxy

import (
	"AdBlockProxy/libs/config"
	"github.com/liuzheng/golog"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"time"
)

const lname = "proxy"

type Server struct {
	listener   net.Listener
	Addr       string
	credential string
}

func (s *Server) Start() {
	//var err error
	golog.Info(lname, "Start service at %v", s.Addr)
	//s.listener, err = net.Listen("tcp", s.Addr)
	//if err != nil {
	//	golog.Error(lname, "Error listening: %v", err)
	//	os.Exit(1)
	//}
	//defer s.listener.Close()
	//for {
	//	conn, err := s.listener.Accept()
	//	if err != nil {
	//		golog.Error(lname, "Error accepting: %v", err)
	//		os.Exit(1)
	//	}
	//
	//	go s.handler(conn)
	//}
	http.HandleFunc("/", s.handler)

	http.ListenAndServe(s.Addr, nil)
}
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	golog.Debug(lname, "%v", r.Host)
	golog.Debug(lname, "%v", r.RequestURI)
	w.Header().Set("proxy", "AdBlockProxy1.0")

	black_action := ""
	// block checker
	if _, ok := config.Config.Blacklist[r.Host]; ok {
		for host, v := range config.Config.Blacklist[r.Host] {
			golog.Debug(lname, "host:%v ", host)
			mached, err := regexp.MatchString(v.UriRe, r.RequestURI)
			if err != nil {
				golog.Error(lname, "%v", err)
			}
			if mached {
				black_action = v.Action
				break
			}
		}
	}

	// block action
	if black_action != "" {
		w.Write([]byte("blocked"))
		return
	}

	// none block do
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	golog.Debug(lname, "%v  %v  %v", r.Method, r.Header.Get("protocol-scheme"), r.Host+r.RequestURI)
	req, err := http.NewRequest(r.Method, r.Header.Get("protocol-scheme")+"://"+r.Host+r.RequestURI, r.Body)
	//req, err := http.NewRequest(r.Method, "http://www.baidu.com"+r.RequestURI, r.Body)
	if err != nil {
		golog.Error(lname, "%v", err)

	}
	r.Header.Del("protocol-scheme") // use nginx to set this things
	r.Header.Del("If-None-Match")
	r.Header.Del("If-Modified-Since")
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			golog.Debug(lname, "r.Header: %s=%s", k, v[0])
			req.Header.Set(k, v[0])
		}
	}

	for _, c := range r.Cookies() {
		req.AddCookie(c)
	}

	resp, err := client.Do(req)
	if err != nil {
		golog.Error(lname, "client.Do %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		golog.Error(lname, "%v", err)
	}
	//golog.Debug(lname, "%v", body)
	if len(resp.Header) > 0 {
		for k, v := range resp.Header {
			w.Header().Set(k, v[0])
		}
	}
	for _, c := range resp.Cookies() {
		http.SetCookie(w, c)
	}
	w.Write(body)
}
func (s *Server) loadBlockList() {
}
