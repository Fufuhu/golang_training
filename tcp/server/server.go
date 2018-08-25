package main

import (
	"fmt"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "3333"
	TYPE = "tcp"
)

func main() {
	l, err := net.Listen(TYPE, HOST+":"+PORT)

	if err != nil {
		fmt.Println("Error listening", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + HOST + ":" + PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}
		handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)

	fmt.Printf("%d, %s", reqLen, string(buf))

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	conn.Write([]byte("Message received."))
	conn.Close()
}
