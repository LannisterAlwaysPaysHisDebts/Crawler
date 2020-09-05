package main

import (
	"Crawler/config"
	"Crawler/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", config.RpcPort)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", conn)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}
