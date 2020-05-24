package mime

import (
	"bytes"
	"fmt"
)

const (
	Marker = "---GOPASS-SECRET-1.0---\n"
)

type Secret struct {
	Header map[string]string
	Body   []byte
}

func Parse(buf []byte) (*Secret, error) {
	if !bytes.HasPrefix([]byte(Marker), buf) {
		return nil, fmt.Errorf("not supported")
	}
	return &Secret{}, nil
}
