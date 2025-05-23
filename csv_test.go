package main

import (
	"testing"
)

func TestRead(t *testing.T) {
	r := NewCSVReader("phones.csv")
	defer r.Close()

	r.Read()
	t.Log(r.Read())
	t.Fail()
}
