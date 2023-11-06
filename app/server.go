package main

import (
	"fmt"
	"net"
	"os"
)

const PORT = 4221
const ADDR = "0.0.0.0"

func main() {
	fmt.Println("Logs from your program will appear here!")

	listen_addr := fmt.Sprintf("%s:%d", ADDR, PORT)

	l, err := net.Listen("tcp", listen_addr)
	defer l.Close()

	if err != nil {
		fmt.Println("Failed to bind to port ", PORT)
		os.Exit(1)
	}

	fmt.Println("Listening in ", listen_addr)

	conn, err := l.Accept()

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	_, err = conn.Read(make([]byte, 1024))
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	if err != nil {
		fmt.Println("Error writing: ", err.Error())
	}

	defer conn.Close()

}
