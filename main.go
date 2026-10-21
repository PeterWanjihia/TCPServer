package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Failed to create TCP listener:", err)
	}

	defer listener.Close()
	fmt.Println("TCP listener created succesfully")
	fmt.Println("Listening on:", listener.Addr())

	fmt.Println("Waiting for a client...")
	conn, err := listener.Accept()

	if err != nil {
		fmt.Println("Failed to accept connection:", err)
		return
	}

	defer conn.Close()

	fmt.Println("Client coonnected:", conn.RemoteAddr())

}
