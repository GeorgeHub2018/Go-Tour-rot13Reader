package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) (r byte) {
	if b >= 'a' && b <= 'z' {
		// Rotate lowercase letters 13 places.
		if b >= 'm' {
			return b - 13
		}
		return b + 13
	} else if b >= 'A' && b <= 'Z' {
		// Rotate uppercase letters 13 places.
		if b >= 'M' {
			return b - 13
		}
		return b + 13
	}
	return b
}

func (r rot13Reader) Read(b []byte) (n int, e error) {
	count, err := r.r.Read(b)
	if count == 0 || err != nil {
		return 0, errors.New("empty read")
	}

	for i := 0; i < count; i++ {
		b[i] = rot13(b[i])
		n++
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
