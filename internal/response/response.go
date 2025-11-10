package response

import (
	"fmt"
	"io"
	"strconv"

	"boot.taran1s/internal/headers"
)

type StatusCode int

const (
	StatusOK                  StatusCode = 200
	StatusBadRequest          StatusCode = 400
	StatusInternalServerError StatusCode = 500
)

var statusName = map[StatusCode]string{
	StatusOK:                  "OK",
	StatusBadRequest:          "Bad Request",
	StatusInternalServerError: "Internal Server Error",
}

func (sc StatusCode) String() string {
	return statusName[sc]
}

func WriteStatusLine(w io.Writer, statusCode StatusCode) error {
	status := string(fmt.Sprintf("HTTP/1.1 %d %s \r\n", statusCode, statusCode))
	_, err := w.Write([]byte(status))
	if err != nil {
		return err
	}
	return nil
}

func GetDefaultHeaders(contentLen int) headers.Headers {
	h := headers.NewHeaders()

	h.Set("Content-Length", strconv.Itoa(contentLen))
	h.Set("Connection", "close")
	h.Set("Content-type", "text/plain")
	return *h
}

func WriteHeaders(w io.Writer, h headers.Headers) error {
	headers := h.GetAll()

	for k, v := range headers {
		_, err := w.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
		if err != nil {
			return err
		}
	}

	_, err := w.Write([]byte("\r\n"))
	if err != nil {
		return err
	}
	return nil
}

func WriteBody(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	if err != nil {
		return err
	}
	return nil
}
