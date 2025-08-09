package main

import (
	"encoding/json"
	"errors"
	"net"
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
	return nil
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
