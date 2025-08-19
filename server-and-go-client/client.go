package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"net/url"
	"os"
	"slices"

	"github.com/gorilla/websocket"
)

type client struct {
	name string
	conn *websocket.Conn
}

type Message struct {
	Name    string
	Message string
}

func NewClient(name string) client {
	client := client{
		name: name,
	}

	return client
}

func (c *client) connect(ipWithPort string) error {
	u := url.URL{Scheme: "ws", Host: ipWithPort, Path: "/chat"}

	fmt.Printf("\nConnecting to: %v\n", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	fmt.Printf("Connection successful\n\n")

	c.conn = conn
	go c.handleIncommingMessages()

	return nil
}

func (c *client) handleIncommingMessages() {
	var msg Message

	for {
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				fmt.Println("Server closed")
				os.Exit(0)
			}
			fmt.Println("hier")
			fmt.Println(err)
		}

		if msg.Name != c.name {
			fmt.Printf("\x1b[38;5;%vm", getColorForName(msg.Name))
			fmt.Printf("%s: %s\n", msg.Name, msg.Message)
			fmt.Print("\x1b[0m")
		}
	}
}

func getColorForName(name string) byte {
	hash := fnv.New32()
	hash.Write([]byte(name))
	ansiColor := byte(hash.Sum32() % 232)

	invalidColors := []byte{0, 16, 17, 52}
	for {
		if !slices.Contains(invalidColors, ansiColor) {
			break
		}
		ansiColor++
	}

	return ansiColor
}

func (c *client) sentMessage(msg string) error {
	if c.conn == nil {
		return errors.New("Client is not connected to the server!")
	}

	data := Message{
		Name:    c.name,
		Message: msg,
	}

	err := c.conn.WriteJSON(data)
	return err
}
