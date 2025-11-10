package headers

import (
	"bytes"
	"fmt"
)

type Headers map[string]string

func NewHeaders() Headers {
	return make(Headers)
}

var MALFORMED_FIELD_LINE = fmt.Errorf("Malformed Field Line")
var MALFORMED_FIELD_NAME = fmt.Errorf("Malformed Field Name")

func parseHeader(fieldLine []byte) (string, string, error) {
	parts := bytes.SplitN(fieldLine, []byte(":"), 2)
	if len(parts) != 2 {
		return "", "", MALFORMED_FIELD_LINE
	}

	if bytes.HasSuffix(parts[0], []byte(" ")) {
		return "", "", MALFORMED_FIELD_NAME
	}

	name := bytes.TrimSpace(parts[0])
	value := bytes.TrimSpace(parts[1])

	return string(name), string(value), nil
}

var SEPARATOR = []byte("\r\n")

func (h Headers) Parse(data []byte) (int, bool, error) {
	// No CLRF means we are awaiting data
	read := 0
	done := false

	for {
		idx := bytes.Index(data[read:], SEPARATOR)

		if idx == -1 {
			break
		}

		if idx == 0 {
			done = true
			read += len(SEPARATOR)
			break
		}

		name, value, err := parseHeader(data[read : read+idx])
		if err != nil {
			return 0, done, err
		}

		read += idx + len(SEPARATOR)
		h[name] = value
	}

	return read, done, nil
}
