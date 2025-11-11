package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"boot.taran1s/internal/request"
	"boot.taran1s/internal/response"
	"boot.taran1s/internal/server"
)

const port = 8888

func get200() []byte {
	return []byte(`
	<html>
  <head>
    <title>200 OK</title>
  </head>
  <body>
    <h1>Success!</h1>
    <p>Your request was an absolute banger.</p>
  </body>
</html>`)
}

func get400() []byte {
	return []byte(`
	<html>
  <head>
    <title>400 Bad Request</title>
  </head>
  <body>
    <h1>Bad Request</h1>
    <p>Your request honestly kinda sucked.</p>
  </body>
</html>`)
}

func get500() []byte {
	return []byte(`
	<html>
  <head>
    <title>500 Internal Server Error</title>
  </head>
  <body>
    <h1>Internal Server Error</h1>
    <p>Okay, you know what? This one is on me.</p>
  </body>
</html>`)
}

func handleRequest(w *response.Writer, req *request.Request) {
	// This is just for testing but probably a better way to implement this?
	h := response.GetDefaultHeaders(0)
	body := get200()
	status := response.StatusOK

	switch req.RequestLine.RequestTarget {
	case "/yourproblem":
		body = get400()
		status = response.StatusBadRequest
	case "/myproblem":
		body = get500()
		status = response.StatusInternalServerError
	}

	w.WriteStatusLine(status)
	h.Replace("Content-Length", fmt.Sprintf("%d", len(body)))
	h.Replace("Content-Type", "text/html")
	w.WriteHeaders(h)
	w.WriteBody(body)
}

func main() {
	server, err := server.Serve(port, handleRequest)
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
