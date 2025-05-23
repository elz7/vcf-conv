package main

import (
	"testing"
)

func TestRead(t *testing.T) {
	r := NewCSVReader("phones.csv")
	defer r.Close()

	p, _ := r.Read()
	t.Log(p)
	t.Fatal()
}
