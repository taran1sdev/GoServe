package server

import (
	"fmt"
	"io"
	"log"
	"net"

	"boot.taran1s/internal/response"
)

type Server struct {
	closed bool
}

func runConnection(s *Server, conn io.ReadWriteCloser) {

	response.WriteStatusLine(conn, response.StatusOK)

	b := []byte("Hello World!")

	h := response.GetDefaultHeaders(len(b))

	response.WriteHeaders(conn, h)
	response.WriteBody(conn, b)

	conn.Close()
}

func runServer(s *Server, listener net.Listener) {

	conn, err := listener.Accept()
	if err != nil || s.closed {
		return
	}

	go runConnection(s, conn)
}

func (s *Server) Close() error {
	s.closed = true
	return nil
}

func (s *Server) listen() {

}

func Serve(port uint16) (*Server, error) {
	server := &Server{closed: false}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Failed to start server on port", err)
	}

	go runServer(server, listener)

	return server, nil
}
