package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type TxtReader struct {
	File    *os.File
	Scanner *bufio.Scanner
}

func NewTxtReader(file string) TxtReader {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	return TxtReader{f, s}
}

func (txt TxtReader) Read() (Person, error) {
	if !txt.Scanner.Scan() {
		return Person{}, io.EOF
	}
	str := txt.Scanner.Text()
	fn, ln := randName(str)
	ph := makeMap(strings.Trim(str, ","), "CELL")
	return Person{FirstName: fn, LastName: ln, Phones: ph}, nil
}

func (r TxtReader) Close() {
	r.File.Close()
}
