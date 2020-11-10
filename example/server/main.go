package main

import (
	"fmt"
	"go-zinx/ziface"
	"go-zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("recv from client : msgId = ", request.GetMsgID(), "data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping..ping..ping"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer()
	s.AddRouter(&PingRouter{})
	s.Serve()
}
