package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
	Send(data []byte) error
	SendBuff(data []byte) error
}

type HandFunc func(*net.TCPConn, []byte, int) error
