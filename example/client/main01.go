package main

import (
	"fmt"
	"go-zinx/znet"
	"io"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return
	}
	for {
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("zinx client test msg")))
		if err != nil {
			fmt.Println("Pack error:", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error ", err)
			return
		}
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error ", err)
			break
		}
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("server unpack msgHead", err)
			break
		}
		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error, ", err)
				return
			}
			fmt.Println("---> Recv Server Msg:ID = ", msg.Id, ",Len = ", msg.DataLen, ",.data = ", string(msg.Data))
		}
		time.Sleep(1 * time.Second)
	}
}
