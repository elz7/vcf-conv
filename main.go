package main

import (
	"fmt"
	"os"
	"io"
	"log"
)

func usage() {
	fmt.Println("vcf-conv <input file> <output file>")
	fmt.Println("    Supported file extensions:")
	fmt.Println("    txt,vcf,csv")
	fmt.Println("    Arguments:")
	fmt.Println("      --sep <char> - a separator in txt file (default is the comma)")
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

	switch get_file_ext(in) {
	case "csv":
		r = NewCSVReader(in)
		defer r.Close()
	case "txt":
		r = NewTxtReader(in)
		defer r.Close()
	}

	switch get_file_ext(out) {
	case "vcf":
		w = NewVCFWriter(out)
		defer w.Close()
	}

	for {
		p, err := r.Read()
		if err == io.EOF {
			break
		}
		log.Print(p)
		w.Write(p)
	}
}

func get_file_ext(file string) string {
	i := len(file) - 1
	for file[i] != '.' {
		i -= 1
	}
	return file[(i + 1):]
}
