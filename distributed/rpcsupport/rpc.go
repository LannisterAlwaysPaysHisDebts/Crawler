package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 根据host注册rpc.Serve
func ServeRpc(host string, server interface{}) error {
	_ = rpc.Register(server)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	log.Printf("Linsting on %s", host)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", conn)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
	return nil
}

// 根据host创建对应的rpc.Client连接
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
