package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

type allClients struct {
	connections []net.Conn
	mu          sync.Mutex
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
		connections: []net.Conn{},
		mu:          sync.Mutex{},
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		clients.mu.Lock()
		clients.connections = append(clients.connections, conn)
		clients.mu.Unlock()

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var msg Message
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		jsonData := scanner.Text()

		err := json.Unmarshal([]byte(jsonData), &msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		handleMessage(msg)
	}
}

func handleMessage(msg Message) {
	fmt.Printf("%v: %v\n", msg.Name, msg.Message)
}
