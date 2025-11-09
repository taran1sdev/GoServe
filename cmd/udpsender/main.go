package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8888")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("error", "error", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")

		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error", "error", err)
		}

		conn.Write([]byte(str))
		if err != nil {
			log.Fatal("error", "error", err)
		}
	}
}
