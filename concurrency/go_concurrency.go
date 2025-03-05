package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	clients = make(map[net.Conn]bool)
	mutex   = sync.Mutex{}
)
var messages = make(chan string)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	defer cleanup()
	fmt.Println("Chat server started on port 8080...")
	go startMessageHandler()
	for {
		conn, err := listener.Accept() // Accept new client connections
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		mutex.Lock()
		clients[conn] = true // Add client to list
		mutex.Unlock()

		go handleClient(conn) // Handle client in a new goroutine
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()

	name := conn.RemoteAddr().String()
	fmt.Println("IP : ", name, "connected",conn)
	
	messages <- name + " joined the chat"

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		// message := name + ": " + scanner.Text()
		// broadcast(message)
		fmt.Println("scanner Text ", scanner.Text())
		messages <- name + ": " + scanner.Text()
	}

	mutex.Lock()
	delete(clients, conn) // Remove client from list
	mutex.Unlock()

	fmt.Println(name, "disconnected")
	messages <- name + " left the chat"
}

func broadcast(message string) {
	mutex.Lock()
	for conn := range clients {
		fmt.Fprintln(conn, message) // Send message to each client
	}
	mutex.Unlock()
}

func startMessageHandler() {
	for message := range messages {
		broadcast(message)
	}
}
func cleanup() {
	mutex.Lock()
	for conn := range clients {
		fmt.Fprintln(conn, "Server shutting down...")
		conn.Close()
	}
	mutex.Unlock()
}
