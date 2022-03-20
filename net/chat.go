package main

// chat.go -> Server -> build and run this file first

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	serverHost = flag.String("h", "localhost", "host")
	serverPort = flag.Int("p", 3090, "port")
)

// Client1 -> Server -> HanldeConnection(Client1)

// HandleConnection handles the client connection of a single user
func HandleConnection(conn net.Conn) {
	defer conn.Close()

	// Create the client channel
	message := make(chan string)
	go MessageWrite(conn, message)

	// Get the client's name
	// for example -> Client1:2560 platzi.com, 38 -> platzi.com:38
	clientName := conn.RemoteAddr().String()

	// Send just to the client his name
	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)
	// Send the message to all the clients
	messages <- fmt.Sprintf("New client is here, name %s\n", clientName)
	// Add the client to the list of clients
	incomingClients <- message

	// Read the messages from the client
	// If the loop breaks, that means that the client has disconnected
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	// Remove the client from the list of clients
	leavingClients <- message
	messages <- fmt.Sprintf("%s said goodbye!", clientName)
}

// MessageWrite receives messages from the channel and writes them to the client
func MessageWrite(conn net.Conn, messages <-chan string) {
	for messsage := range messages {
		fmt.Fprintln(conn, messsage)
	}
}

// Broadcast sends the message to all the clients, and handles incoming and outgoing connections
func Broadcast() {
	// Map the clients to a boolean
	clients := make(map[Client]bool)

	for {
		// Multiplex the messages
		select {
		case message := <-messages: // Send the message to all the clients
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients: // We get a new client
			clients[newClient] = true
		case leavingClient := <-leavingClients: // A client has disconnected
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	flag.Parse()

	// Create the server and listen to it
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *serverHost, *serverPort))
	if err != nil {
		log.Fatal(err)
	}

	// Start the broadcast
	go Broadcast()

	// Listen for connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conn)
	}
}
