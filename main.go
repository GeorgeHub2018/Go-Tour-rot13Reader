package main

import (
	"errors"
	"io"
	"os"
	"strings"
)

const (
	input  string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!"
	output string = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm!"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(str byte) byte {
	index := strings.Index(input, string(str))
	if index > -1 {
		return byte(output[index])
	}
	return 0
}

func (r rot13Reader) Read(b []byte) (n int, e error) {
	rotByte := make([]byte, len(b))
	bi, err := r.r.Read(rotByte)
	if bi == 0 || err != nil {
		return 0, errors.New("empty read")
	}

	n = 0
	for i := 0; i < bi; i++ {
		b[i] = rot13(rotByte[i])
		n++
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
