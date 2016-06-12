package main

/*
  connect to echo server with
    telnet localhost 6000
*/

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(os.Stderr, err)
			continue
		}
        // handle in goroutine, a lightweight thread of execution
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	buf := make([]byte, 4096)

	for {
		n, err := c.Read(buf)
		if err != nil || n == 0 {
			c.Close()
			break
		}
		n, err = c.Write(buf[0:n])
		if err != nil {
			c.Close()
			break
		}
	}
	fmt.Printf("Connection from %v closed.\n", c.RemoteAddr())
}
