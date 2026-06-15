package main

import (
	"bufio"
	"fmt"
	"net"
)


func handleConnection(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	fmt.Printf("Connected with %v\n", addr)
	conn.Write([]byte("Hello to " + addr + "!\n"))
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
LOOP:
	for scanner.Scan() {
		text := scanner.Text()
		switch {
		case text == "Exit":
			conn.Write([]byte("Bye!\n"))
			break LOOP
		case text != "":
			conn.Write([]byte("You entered " + text + "\n"))
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}


func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}