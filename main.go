package main

import (
	"fmt"
	"net"

	"os"
)

func main() {

	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)

	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()
	readMessage := make(chan string)
	writeMessage := make(chan string)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
