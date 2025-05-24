package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

type VCFWriter struct {
	File *os.File
}

func NewVCFWriter(file string) VCFWriter {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	return VCFWriter{
		File: f,
	}
}

func (w VCFWriter) Write(p Person) {
	f := bufio.NewWriter(w.File)
	defer f.Flush()

	write := func(str string, args ...any) {
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
		f.WriteString(
			fmt.Sprintf(str, args...))
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
	write("PHOTO;ENCODING=BASE64;%v\n", p.Photo)
	write("END:VCARD\n")
}

func (w VCFWriter) Close() {
	w.File.Close()
}
