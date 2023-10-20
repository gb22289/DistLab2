package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
	}
}

func write(conn net.Conn) {
	for {
		stdin := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter message:")
		msg, _ := stdin.ReadString('\n')
		fmt.Print(msg)
		fmt.Fprint(conn, msg)
	}
	//TODO Continually get input from the user and send messages to the server.
}

func main() {
	// Get the server address and port from the commandline arguments.
	server := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	conn, _ := net.Dial("tcp", *server)
	go read(conn)
	write(conn)
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
