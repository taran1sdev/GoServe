package main

import (
	"fmt"
	"log"
	"net"

	"boot.taran1s/internal/request"
)

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}

		r, err := request.RequestFromReader(conn)
		if err != nil {
			log.Fatal("error", "error", err)
		}

		fmt.Printf("Request line:\n- Method: %s\n- Target: %s\n- Version: %s\n", r.RequestLine.Method, r.RequestLine.RequestTarget, r.RequestLine.HttpVersion)

		fmt.Printf("Headers:")

		for k, v := range r.Headers.GetAll() {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
}
