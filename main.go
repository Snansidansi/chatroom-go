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
		println("server")
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
}
