package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"net"
	"slices"
)

type client struct {
	name string
	conn net.Conn
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
	conn, err := net.Dial("tcp", ipWithPort)
	if err != nil {
		return err
	}

	c.conn = conn
	go c.handleIncommingMessages()

	return nil
}

func (c *client) handleIncommingMessages() {
	scanner := bufio.NewScanner(c.conn)
	var msg Message

	for scanner.Scan() {
		data := scanner.Bytes()

		err := json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println(err)
			continue
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	jsonData = append(jsonData, '\n')

	_, err = c.conn.Write(jsonData)
	return err
}
