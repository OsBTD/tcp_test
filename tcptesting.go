package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatal("error listening", err)
	}
	defer listener.Close()
	fmt.Println("Server Started")
	var clients []net.Conn
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("error accepting connection", err)
		}
		clients = append(clients, conn)
		go handleclient(conn)
		fmt.Println(clients)

	}
}

func handleclient(conn net.Conn) {
	read := bufio.NewReader(conn)
	for {
		msg, err := read.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("client disconnected")
				return

			} else {
				log.Println("error reading", err)
				return

			}
		}
		fmt.Printf("received message : %s", msg)

		_, erro := conn.Write([]byte("Message received " + msg))
		if erro != nil {
			log.Println("error receiving", err)
		}

	}
}
