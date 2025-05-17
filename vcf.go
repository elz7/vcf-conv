package main

import (
	"os"
	"fmt"
	"bufio"
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

func formatPhone(p string) string {
	return p[:3] + "-" + p[3:6] + "-" + p[6:]
}

func (w VCFWriter) Write(record Person) {
	f := bufio.NewWriter(w.File)
	defer f.Flush()

	f.WriteString("BEGIN:VCARD\n")
	f.WriteString("VERSION:2.1\n")
	f.WriteString(fmt.Sprintf("N:%v;%v;;;\n", record.LastName, record.FirstName))
	f.WriteString(fmt.Sprintf("FN:%v %v\n", record.FirstName, record.LastName))
	for p, t := range record.Phones {
		f.WriteString(fmt.Sprintf("TEL;%v:%v\n", t, formatPhone(p)))
	}
	for e, t := range record.Emails {
		f.WriteString(fmt.Sprintf("EMAIL;%v:%v\n", t, e))
	}
	//write("NOTE:%v", record.Note)
	//write("PHOTO;ENCODING=BASE64;%v", record.Photo)
	f.WriteString("END:VCARD\n")
}

func (w VCFWriter) Close() {
	w.File.Close()
}
