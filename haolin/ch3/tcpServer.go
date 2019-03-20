package main

import "net"

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		panic("listen error")
	}
	conn, err := listener.Accept()


}
