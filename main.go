package main

import (
	"fmt"
	"os"
	"bytes"
	"io"
)

type Person struct {
	FirstName string
	LastName  string
	Phones    map[string]string
	Emails    map[string]string
	Note      string
	Photo     string
}

func (p Person) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("[%v:", p.FirstName + " " + p.LastName))
	for p, _ := range p.Phones {
		buffer.WriteString(fmt.Sprintf(" %v", p))
	}
	buffer.WriteString("]\n")
	return buffer.String()
}

type Writer interface {
	Write(rec Person)
	Close()
}

type Reader interface {
	Read() (Person, error)
	Close()
}

func usage() {
	fmt.Println("vcf-conv <input file> <output file>")
	fmt.Println("    Supported file extensions:")
	fmt.Println("    vcf,csv")
	// fmt.Println("    ")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	in := os.Args[1]
	out := os.Args[2]

	var r Reader
	var w Writer

	switch getFileExt(in) {
	case "csv":
		r = NewCSVReader(in)
		defer r.Close()
	}

	switch getFileExt(out) {
	case "vcf":
		w = NewVCFWriter(out)
		defer w.Close()
	}

	for {
		p, err := r.Read()
		if err == io.EOF {
			break
		}

		for k, _ := range p.Phones {
			p.FirstName = k
			break
		}

		w.Write(p)
	}
}

func getFileExt(file string) string {
	i := len(file) - 1
	for file[i] != '.' {
		i -= 1
	}
	return file[(i + 1):]
}
