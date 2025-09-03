package funcs

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"unicode"
)

func HandleClient(conn net.Conn) { // net.Conn = the network connection between the server and the client
	defer conn.Close()

	// Check max client limit
	clientsMutex.Lock()
	if len(clients) > 10 {
		clientsMutex.Unlock()
		fmt.Fprintln(conn, "Server is full. Try again later.")
		return
	}
	clientsMutex.Unlock()

	var clientName string

	fmt.Fprintln(conn, LinuxAscii)
	for {
		fmt.Fprint(conn, "[ENTER YOUR NAME]:")

		nameScanner := bufio.NewScanner(conn)
		if !nameScanner.Scan() { // Waits for user input
			return
		}

		clientName = nameScanner.Text()

		if !isValidName(clientName) {
			fmt.Fprintln(conn, "Invalid name. Please use only printable characters, and max of 30 characters.")
			continue
		}

		if isExist(clientName) {
			fmt.Fprintln(conn, "Client name exist, please chose diffrent name")
			continue
		}

		break
	}

	client := &Client{
		Name: clientName,
		Conn: conn,
	}

	// Add client to map
	clientsMutex.Lock()
	clients[conn] = client
	clientsMutex.Unlock()

	welcomeMsg := fmt.Sprintf("%s has joined our chat...", client.Name)
	fmt.Println(welcomeMsg)   // Server log
	chanMessage <- welcomeMsg // For sending to all clients

	// Send chat history to the new client
	clientsMutex.Lock()
	for _, msg := range serverLogHistory {
		fmt.Fprintln(conn, msg)
	}
	clientsMutex.Unlock()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // Runs as long as the client keeps sending data

		if scanner.Text() == "" {
			continue
		}

		client.message = scanner.Text()

		if !isValidInput(client.message) {
			fmt.Fprintln(conn, "ERROR: mesage is too long or have invalid caracters!, max lenght is 2000 caracter")
			continue
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		clientMessage := fmt.Sprintf("[%s] [%s]: %s", timestamp, client.Name, client.message)

		fmt.Println(clientMessage)   // Server log
		chanMessage <- clientMessage // For sending to all clients
	}

	// Client disconnected
	clientsMutex.Lock()
	delete(clients, conn)
	clientsMutex.Unlock()

	leaveMsg := fmt.Sprintf("%s has left the chat", client.Name)
	fmt.Println(leaveMsg)
	chanMessage <- leaveMsg

}

func isExist(str string) bool {
	for _, client := range clients {
		if client.Name == str {
			return true
		}
	}
	return false
}

func isValidName(name string) bool {
	if len(name) == 0 || len(name) > 30 {
		return false
	}
	for _, r := range name {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

func isValidInput(name string) bool {
	if len(name) > 2000 {
		return false
	}
	for _, r := range name {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}
