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
	message := "Hello, World!"

	tcpAddr, err := net.ResolveTCPAddr("tcp", HOST+":"+PORT)

	if err != nil {
		fmt.Println("ResolveTCPAddr failed.", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpAddr)
	if err != nil {
		fmt.Println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Writer to server:", message)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)

	if err != nil {
		fmt.Println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Reply from server=", string(reply))
}
