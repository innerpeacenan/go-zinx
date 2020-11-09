package main

import "go-zinx/znet"

func main() {
	s := znet.NewServer("[zinx v2.0]")
	s.Serve()

}
