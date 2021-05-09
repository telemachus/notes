package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening via tcp on :9999")

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %q\n", msg)
	c.Close()
}

func client(msg string) {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sending: %q\n", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		log.Fatal(err)
	}

	c.Close()
}

func main() {
	go server()
	for {
		var input string
		fmt.Scanln(&input)

		if input == "quit" {
			break
		}

		go client(input)
	}
}
