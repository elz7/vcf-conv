package main

import "errors"

var ErrEmptyLine = errors.New("empty line")

type Person struct {
	FirstName string
	LastName  string
	Phones    map[string]string
	Emails    map[string]string
	Note      string
}

type Reader interface {
	Read() (Person, error)
	Close() error
}

type Writer interface {
	Write(p Person)
	Close() error
}

type CommandLine struct {
	InputFile      string
	OutputFile     string
	PrintCSVFields bool
}
