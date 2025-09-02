package funcs

import (
	"fmt"
	"net"
)

func ServerHandler(port string) {

	// Accepts incoming TCP connections
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer server.Close()

	fmt.Println("Server is listening on port", port)

	go Broadcast() // Start listening to new messages

	for {
		conn, err := server.Accept() // Wait for a new clients
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go HandleClient(conn) // Start a new goroutine for this client
	}
}

func Broadcast() {
	for msg := range chanMessage { // A loop that waits for new messages on the channel
		// Save message to history
		serverLogHistory = append(serverLogHistory, msg)

		// Send to all clients
		clientsMutex.Lock()
		for _, client := range clients {
			fmt.Fprintln(client.Conn, msg)
		}
		clientsMutex.Unlock()
	}
}
