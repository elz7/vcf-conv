package main

import (
	"bufio"
	"os"
)

type TxtReader struct {
	File   *os.File
	Reader *bufio.Reader
}

func NewTxtReader(file string) TxtReader {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	return TxtReader{
		File:   f,
		Reader: r,
	}
}

func (r TxtReader) Read() (Person, error) {
	str, err := r.Reader.ReadString(byte(','))
	if err != nil {
		return Person{}, err
	}
	phones := make(map[string]string)
	phones[str[:(len(str)-2)]] = "CELL"
	return Person{
		FirstName: str[:(len(str) - 2)],
		Phones:    phones,
	}, nil
}

func (r TxtReader) Close() {
	r.File.Close()
}
