package main

import (
	"encoding/csv"
	"os"
)

type CSVReader struct {
	File   *os.File
	Reader *csv.Reader
}

func NewCSVReader(file string) CSVReader {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)
	return CSVReader{
		File:   f,
		Reader: r,
	}
}

func (r CSVReader) Read() (Person, error) {
	rec, err := r.Reader.Read()
	if err != nil {
		return Person{}, err
	}

	phones := make(map[string]string)
	phones[rec[0]] = "CELL"
	return Person{
		Phones: phones,
	}, nil
}

func (r CSVReader) Close() {
	r.File.Close()
}
