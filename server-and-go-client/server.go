package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type allClients struct {
	connections map[uint]*websocket.Conn
	mu          sync.Mutex
	index       uint
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func StartServer(port string) {
	clients := allClients{
		mu:          sync.Mutex{},
		connections: map[uint]*websocket.Conn{},
		index:       0,
	}

	if port[0] != ':' {
		port = ":" + port
	}

	fmt.Printf("Server is listening on port %s\n", port)

	http.HandleFunc("/chat", clients.chatWebSocketHandler)
	http.ListenAndServe(port, nil)
}

func (ac *allClients) chatWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	connIndex := ac.index
	ac.index++

	ac.mu.Lock()
	ac.connections[connIndex] = conn
	ac.mu.Unlock()

	defer func() {
		ac.mu.Lock()
		delete(ac.connections, connIndex)
		ac.mu.Unlock()

		conn.Close()
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				fmt.Println("Client disconnected")
				return
			}
			fmt.Println(err)
			return
		}

		handleMessage(message, messageType, ac)
	}
}

func handleMessage(msg []byte, messageType int, clients *allClients) {
	clients.mu.Lock()
	for _, conn := range clients.connections {
		if err := conn.WriteMessage(messageType, msg); err != nil {
			fmt.Println(err)
			continue
		}
	}
	clients.mu.Unlock()
}
