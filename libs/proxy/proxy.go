package proxy

import (
	"crypto/tls"
	//"AdBlockProxy/libs/config"
	"github.com/liuzheng/golog"
	"io/ioutil"
	"libs/config"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

const lname = "proxy"

func Start() {
	//var err error
	addr := config.Config.Server + ":" + strconv.Itoa(int(config.Config.Listen))
	golog.Info(lname, "Start service at %v", addr)
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
	http.HandleFunc("/", handler)

	http.ListenAndServe(addr, nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	golog.Debug(lname, "%v", r.Host)
	golog.Debug(lname, "%v", r.RequestURI)
	w.Header().Set("proxy", "AdBlockProxy1.0")

	var matchURI config.URI
	var client *http.Client
	// block checker
	if _, ok := config.Config.Blacklist[r.Host]; ok {
		for host, v := range config.Config.Blacklist[r.Host] {
			golog.Debug(lname, "host:%v ", host)
			mached, err := regexp.MatchString(v.UriRe, r.RequestURI)
			if err != nil {
				golog.Error(lname, "%v", err)
			}
			if mached {
				matchURI = v
				break
			}
		}
	}

	// block action
	switch matchURI.Action {
	case "proxy":
		proxy, _ := url.Parse(matchURI.Proxy)
		proxy.User = url.UserPassword(config.Config.ProxyList.Username, config.Config.ProxyList.Password)
		tr := &http.Transport{
			Proxy:           http.ProxyURL(proxy),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{
			Transport: tr,
			Timeout:   5 * time.Second,
		}
	case "":
		client = &http.Client{
			Timeout: 5 * time.Second,
		}
	case "block":
		if matchURI.Code == 0 {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(matchURI.Code)
		}
		w.Write([]byte("blocked"))
		return
	default:
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("blocked"))
		return
	}
	//if black_action != "" {
	//	w.WriteHeader(http.StatusForbidden)
	//	w.Write([]byte("blocked"))
	//	return
	//}

	// none block do
	protocol_scheme := r.Header.Get("protocol-scheme")
	golog.Debug(lname, "%v  %v  %v", r.Method, protocol_scheme, r.Host+r.RequestURI)
	if protocol_scheme == "" {
		golog.Error(lname, "Proto Scheme is not defined")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error, Please check the log"))
		return
	}
	req, err := http.NewRequest(r.Method, protocol_scheme+"://"+r.Host+r.RequestURI, r.Body)
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
func loadBlockList() {
}
