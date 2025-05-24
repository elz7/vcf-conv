package main

import (
	"fmt"
	"os"
	"io"
	"log"
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
	if len(os.Args) < 2 {
		usage()
	}

	if _, ok := get_arg_value("--print-csv-fields", os.Args); ok {
		fmt.Println("firstname,lastname,phones,emails,note,photo")
		os.Exit(0)
	}

	in := os.Args[1]
	out := get_file_name(in) + ".vcf"
	if v, ok := get_arg_value("--out", os.Args); ok {
		out = v
	}

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

func get_file_name(file string) string {
	i := len(file) - 1
	for file[i] != '.' {
		i -=1
	}
	return file[:i]
}

func get_arg_value(arg string, args []string) (string, bool) {
	for i, a := range args {
		if a == arg {
			if ok := i+1 < len(args); ok {
				return args[i+1], ok
			} else {
				return "", true
			}
		}
	}
	return "", false
}
