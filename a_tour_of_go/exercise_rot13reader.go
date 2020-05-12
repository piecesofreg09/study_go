package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot(b byte) byte {
	switch {
	case b >= 65 && b <= 77:
		return b + 13
	case b >= 78 && b <= 90:
		return b - 78 + 65;
	case b >=97 && b <= 109:
		return b + 13
	case b >= 110 && b <= 122:
		return b - 110 + 97
	default:
		return b
	}
	
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	for i := range b {
		b[i] = rot(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
