package main

import (
	"bufio"
	"errors"
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	defer c.Close()

	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		input := scanner.Text()

		if input == "PING" {
			c.Write([]byte("+PONG\r\n"))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("connection error: %v", err)
	}
}

func main() {
	l, err := net.Listen("tcp", ":6380")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("cache-it listening on :6380")

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			// a closed listener never recovers, so continuing here would busy loop
			if errors.Is(err, net.ErrClosed) {
				log.Println("listener closed, shutting down")
				return
			}
			log.Println("accept:", err)
			continue
		}

		go handleConnection(conn)
	}
}
