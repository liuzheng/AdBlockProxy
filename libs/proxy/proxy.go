package proxy

import (
	"github.com/liuzheng/golog"
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
}
