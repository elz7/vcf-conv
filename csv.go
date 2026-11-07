package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVReader struct {
	file *os.File
	csv  *csv.Reader
}

func NewCSVReader(file string) (Reader, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return CSVReader{file: f, csv: csv.NewReader(f)}, nil
}

func (r CSVReader) Read() (Person, error) {
	rec, err := r.csv.Read()
	if err != nil {
		return Person{}, err
	}

	fmt.Println(len(rec))

	var person Person

	if phone, ok := getOrNone(rec, 0); ok {
		person.Phones = makeMap(phone, "CELL")
	}

	if email, ok := getOrNone(rec, 1); ok {
		person.Emails = makeMap(email, "HOME")
	}

	if firstname, ok := getOrNone(rec, 2); ok {
		person.LastName = firstname
	}

	if lastname, ok := getOrNone(rec, 3); ok {
		person.LastName = lastname
	}

	if note, ok := getOrNone(rec, 4); ok {
		person.Note = note
	}

	if person.FirstName == "" && person.LastName == "" {
		firstname, lastname := randName(person.Phones["CELL"])
		if person.FirstName == "" {
			person.FirstName = firstname
		}
		if person.LastName == "" {
			person.LastName = lastname
		}
	}

	return person, nil
}

func (r CSVReader) Close() error {
	return r.file.Close()
}
