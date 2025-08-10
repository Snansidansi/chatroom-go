package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
)

type allClients struct {
	connections map[uint]net.Conn
	mu          sync.Mutex
	index       uint
}

func StartServer(ip, port string) error {
	fullAddress := ip + ":" + port
	listener, err := net.Listen("tcp", fullAddress)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Printf("Server is listening on: %s\n", fullAddress)

	clients := allClients{
		mu:          sync.Mutex{},
		connections: map[uint]net.Conn{},
		index:       0,
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleClient(conn, &clients)
	}
}

func handleClient(conn net.Conn, clients *allClients) {
	connIndex := clients.index
	clients.index++

	clients.mu.Lock()
	clients.connections[connIndex] = conn
	clients.mu.Unlock()

	defer func() {
		clients.mu.Lock()
		delete(clients.connections, connIndex)
		clients.mu.Unlock()

		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return
			}

			fmt.Println(err)
			continue
		}

		handleMessage(msg, clients)
	}
}

func handleMessage(msg []byte, clients *allClients) {
	clients.mu.Lock()
	for _, conn := range clients.connections {
		conn.Write(msg)
	}
	clients.mu.Unlock()
}
