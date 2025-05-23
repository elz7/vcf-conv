package main

import (
	"os"
	"encoding/csv"
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

	fn := rec[0]
	ln := rec[1]
	ph := make_map(rec[2], "CELL")
	em := make_map(rec[3], "HOME")
	nt := rec[4]
	pt := rec[5]

	return Person{fn, ln, ph, em, nt, pt}, nil
}

func (r CSVReader) Close() {
	r.File.Close()
}
