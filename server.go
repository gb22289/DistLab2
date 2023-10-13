package main

import (
	"bufio"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	fmt.Println(err.Error())
	panic(err)
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	for {
		conn, err := ln.Accept()
		if err != nil {
			handleError(err)
		}
		conns <- conn
	}
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	reader := bufio.NewReader(client)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			handleError(err)
		}
		fmt.Println(msg)
		fmt.Println("Ok...")
		message := Message{clientid, msg}
		msgs <- message
	}
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
}

func main() {
	//TODO Create a Listener for TCP connections on the port given above.
	ln, err := net.Listen("tcp", ":8030")
	if err != nil {
		handleError(err)
	}
	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	clientNumber := 0
	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			clients[clientNumber] = conn
			go handleClient(conn, clientNumber, msgs)
			clientNumber++
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
		case msg := <-msgs:
			for i := 0; i < len(clients); i++ {
				if i != msg.sender {
					fmt.Fprint(clients[i], msg.message)
				}
			}
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
		}

	}
}
