package main

import (
	"time"
)

type Client struct {
	name              string
	joinTime          time.Time
	ipWithPort        string
	countMessagesSent int
}

func NewClient(name string) Client {
	client := Client{
		name: name,
	}

	return client
}

func (c *Client) connect(ipWithPort string) error {
	c.ipWithPort = ipWithPort
	c.joinTime = time.Now()

	// Connect to server
	return nil
}
