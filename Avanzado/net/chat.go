package main

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
	leaveClients    = make(chan Client)
	messages        = make(chan string)
)
var (
	host = flag.String("h", "localhost", "hostname")
	port = flag.Int("p", 3090, "port")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWrite(conn, message)
	clientName := conn.RemoteAddr().String()
	message <- fmt.Sprintf("Welcome to the server, %s\n", clientName)
	message <- fmt.Sprintf("New client is here,name %s\n", clientName)
	incomingClients <- message
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s", clientName, inputMessage.Text())
	}
	leaveClients <- message
	messages <- fmt.Sprintf("%s said goodbye", clientName)
}
func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		fmt.Fprintln(conn, msg)
	}
}
func BroadCast() {
	clients := make(map[Client]bool)
	for {
		select {
		case client := <-incomingClients:
			clients[client] = true
		case client := <-leaveClients:
			delete(clients, client)
			close(leaveClients)
		case msg := <-messages:
			for client := range clients {
				client <- msg
			}
		}
	}
}
func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	go BroadCast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conn)
	}
}
