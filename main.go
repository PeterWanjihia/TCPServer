package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)

		if n > 0 {
			fmt.Printf("[%s] Received: %q\n", conn.RemoteAddr(), buffer[:n])

			written, writeErr := conn.Write(buffer[:n])
			if writeErr != nil {
				fmt.Println("Write failed:", writeErr)
				return
			}
			fmt.Printf("Echoed %d bytes\n", written)
		}
		if err != nil {
			fmt.Println("Read stopped:", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp4", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("Failed to create TCP listener:", err)
		return
	}

	defer listener.Close()
	fmt.Println("TCP listener created succesfully")
	fmt.Println("Listening on:", listener.Addr())
	fmt.Println("Waiting for a client...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			return
		}
		go handleConnection(conn)

	}

}
