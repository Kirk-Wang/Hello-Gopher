package main

import (
	"github.com/Kirk-Wang/Hello-Gopher/17.3/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// telnet localhost 1234
// {"method": "abc.def"}
// {"method": "DemoService.Div", "params":[{"A": 3, "B": 4}], "id": 1}
// {"method": "DemoService.Div", "params":[{"A": 3, "B": 0}], "id": 1234}
func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	// 一直监听
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		// 直接开个协程去处理这个连接
		go jsonrpc.ServeConn(conn)
	}
}
