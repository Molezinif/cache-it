package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

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
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			scanner := bufio.NewScanner(c)

			for scanner.Scan() {
				input := scanner.Text()

				if input == "PING" {
					c.Write([]byte("PONG\n"))
				}
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Erro na conexão: %v\n", err)
			}
		}(conn)
	}
}
