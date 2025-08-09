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

	if len(args) == 2 && args[1] == "server" {
		runServer()
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

func runServer() {
	fmt.Print("Please enter the server ip (nothing for 0.0.0.0): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	ip := scanner.Text()
	if ip == "" {
		ip = "0.0.0.0"
	}

	fmt.Print("Please enter a port: ")
	scanner.Scan()
	port := scanner.Text()

	if err := StartServer(ip, port); err != nil {
		fmt.Println(err)
	}
}
