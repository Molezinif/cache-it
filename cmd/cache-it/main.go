package main

import (
	"bufio"
	"fmt"
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
		fmt.Printf("Erro na conexão: %v\n", err)
	}
}

func main() {
	l, err := net.Listen("tcp", ":6380")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Servidor TCP rodando na porta 6380...")

	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleConnection(conn)
	}
}
