package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	fmt.Println("Starting server...")
	listener, err := net.Listen("tcp", ":50000")
	if err != nil {
		fmt.Errorf("[Error] %v", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("[Error] %v", err)
			continue
		}

		go handleConnection(conn)

	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("[info] client disconnected")
			} else {
				fmt.Errorf("[Error] %v", err)
			}
			return
		}

		_, err = conn.Write(data)
		if err != nil {
			fmt.Errorf("[Error] %v", err)
			return
		}
	}
}
