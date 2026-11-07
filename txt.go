package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type TxtReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewTxtReader(file string) (Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return TxtReader{file: f, scanner: bufio.NewScanner(f)}, nil
}

func (txt TxtReader) Read() (Person, error) {
	if !txt.scanner.Scan() {
		return Person{}, io.EOF
	}

	line := txt.scanner.Text()
	if line == "" {
		return Person{}, ErrEmptyLine
	}

	fn, ln := randName(line)
	ph := makeMap(strings.Trim(line, ","), "CELL")

	return Person{FirstName: fn, LastName: ln, Phones: ph}, nil
}

func (r TxtReader) Close() error {
	return r.file.Close()
}
