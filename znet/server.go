package znet

import (
	"fmt"
	"go-zinx/conf"
	"go-zinx/ziface"
	"net"
	"time"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       int
	msgHandler ziface.IMsgHandle
}

func NewServer() ziface.IServer {
	s := &Server{
		Name:       conf.ConfigInstance.Name,
		IPVersion:  "tcp4",
		IP:         conf.ConfigInstance.Host,
		Port:       conf.ConfigInstance.TcpPort,
		msgHandler: NewMsgHandle(),
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.msgHandler.StartWorkerPool()
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")

		var cid uint32
		cid = 0

		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			dealConn := NewConntion(conn, cid, s.msgHandler)
			cid++

			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	for {
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
}
