package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		runClient()
		return
	}

	if len(args) != 3 && args[1] == "server" {
		fmt.Println("Usage: server <port>")
		return
	}

	if len(args) == 3 && args[1] == "server" {
		runServer(args[2])
		return
	}

	println("Invalid arguments. Enter server or nothing (for client).")
}

func runClient() {
	fmt.Print("Please enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	client := NewClient(name)

	fmt.Print("Enter the ip with port of the server: ")
	scanner.Scan()
	ipWithPort := scanner.Text()

	err := client.connect(ipWithPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for scanner.Scan() {
		err := client.sentMessage(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func runServer(port string) {
	StartServer(port)
}
