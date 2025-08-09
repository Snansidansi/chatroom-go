package main

import (
	"fmt"
	"net"
)

type serverConfig struct {
	messages             chan string
	connectedClientsChan chan int
	connectedClients     int
}

func StartServer(ip, port string) error {
	fullAddress := ip + ":" + port

	serverCfg := serverConfig{
		messages:             make(chan string),
		connectedClientsChan: make(chan int),
		connectedClients:     0,
	}

	listener, err := net.Listen("tcp", fullAddress)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Printf("Server is listening on: %s\n", fullAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleClient(conn, &serverCfg)
	}
}

func handleClient(conn net.Conn, serverCfg *serverConfig) {
	defer conn.Close()
}
