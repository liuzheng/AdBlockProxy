package proxy

import (
	"github.com/liuzheng/golog"
	"io/ioutil"
	"net"
	"net/http"
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

	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.Host+r.RequestURI, r.Body)
	//req, err := http.NewRequest(r.Method, "http://www.baidu.com"+r.RequestURI, r.Body)
	if err != nil {
		golog.Error(lname, "%v", err)

	}
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			//fmt.Printf("%s=%s\n", k, v[0])
			req.Header.Set(k, v[0])
		}
	}

	for _, c := range r.Cookies() {
		req.AddCookie(c)
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		golog.Error(lname, "%v", err)
	}
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
