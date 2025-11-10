package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"boot.taran1s/internal/server"
)

const port = 8888

func main() {
	server, err := server.Serve(port)
	if err != nil {
		log.Fatal("Error starting server: %v", err)
	}
	defer server.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server stopped")
}
