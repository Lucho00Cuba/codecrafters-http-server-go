package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"regexp"

	utils "github.com/codecrafters-io/http-server-starter-go/app/http"
)

const (
	PORT         = 4221
	ADDR         = "0.0.0.0"
	AllowedPaths = `^/echo/.*$|^/$`
)

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

	handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	input := make([]byte, 1024)
	_, err := conn.Read(input)
	if err != nil {
		fmt.Println("error reading connection: ", err.Error())
		os.Exit(1)
	}
	req, err := utils.ParseRequest(input)
	if err != nil {
		fmt.Println("error parsing request: ", err.Error())
	}

	statusCode := validatePath(req.Path)
	resp := utils.NewResponse(req, statusCode)
	resp.WriteResponse(conn)
}

func validatePath(path string) int {
	validPathRegex := regexp.MustCompile(AllowedPaths)
	if validPathRegex.MatchString(path) {
		return http.StatusOK
	} else {
		return http.StatusNotFound
	}
}
