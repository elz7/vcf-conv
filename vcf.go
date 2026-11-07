package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type VCFWriter struct {
	file *os.File
}

func NewVCFWriter(file string) (Writer, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return VCFWriter{f}, nil
}

func (w VCFWriter) Write(p Person) {
	f := bufio.NewWriter(w.file)
	defer f.Flush()

	write := func(format string, args ...any) {
		proceed := false
		for _, a := range args {
			if !reflect.ValueOf(a).IsZero() {
				proceed = true
				break
			}
		}
		if !proceed && len(args) != 0 {
			return
		}
		fmt.Fprintf(f, format, args...)
	}

	write("BEGIN:VCARD\n")
	write("VERSION:2.1\n")
	write("N:%v;%v;;;\n", p.LastName, p.FirstName)
	write("FN:%v %v\n", p.FirstName, p.LastName)
	for ph, tp := range p.Phones {
		write("TEL;%v:%v\n", tp, ph)
	}
	for em, tp := range p.Emails {
		write("EMAIL;%v:%v\n", tp, em)
	}
	write("NOTE:%v\n", p.Note)
	write("END:VCARD\n")
}

func (w VCFWriter) Close() error {
	return w.file.Close()
}
