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
		log.Printf("Erro na conexão: %v\n", err)
	}
}

func main() {
	l, err := net.Listen("tcp", ":6380")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Servidor TCP rodando na porta 6380...")

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Println("Server closed!!")
				return // close server
			}
			log.Println("accept:", err)
			continue // continue iteration for new connections
		}

		go handleConnection(conn)
	}
}
