package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Println("vcf-conv <file>")
	fmt.Println("    Supported file extensions:")
	fmt.Println("    txt,csv")
	fmt.Println("    Arguments:")
	fmt.Println("        --out                 output file")
	fmt.Println("        --print-csv-fields    prints how should look csv file")
	os.Exit(0)
}

func main() {
	out := flag.String("out", "", "output file")
	printCSVFields := flag.Bool("print-csv-fields", false, "prints how should look like csv file")
	flag.Parse()
	in := flag.Arg(0)

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "usage: vcf-conf [flags] <file>\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *printCSVFields {
		fmt.Println("firstname,lastname,phones,emails,note,photo")
		os.Exit(0)
	}

	if *out == "" {
		*out = fileNameWithoutExtension(in) + ".vcf"
	}

	var r Reader
	var w Writer

	switch filepath.Ext(in) {
	case ".csv":
		r = NewCSVReader(in)
		defer r.Close()
	case ".txt":
		r = NewTxtReader(in)
		defer r.Close()
	default:
		panic("Unknown file extension")
	}

	switch filepath.Ext(*out) {
	case ".vcf":
		w = NewVCFWriter(*out)
		defer w.Close()
	default:
		panic("Unknown file extension")
	}

	for {
		p, err := r.Read()
		if err == io.EOF {
			break
		}
		log.Print(p)
		w.Write(&p)
	}
}

func fileNameWithoutExtension(file string) string {
	i := len(file) - 1
	for file[i] != '.' {
		i -= 1
	}
	return file[:i]
}
