package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func usage() {
	fmt.Println("vcf-conv <file>")
	fmt.Println()
	fmt.Println("Supported file extensions:")
	fmt.Println("    txt,csv")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("        --out                 output file")
	fmt.Println("        --print-csv-fields    prints how should look csv file")
}

func parseCommandLine() CommandLine {
	flag.Usage = usage
	outputFile := flag.String("out", "", "output file")
	printCSVFields := flag.Bool("print-csv-fields", false, "prints how should look csv file")
	flag.Parse()

	if *printCSVFields {
		fmt.Println("phone,email,firstname,lastname,note")
		fmt.Println("* do not add those to your csv file")
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	inputFile := flag.Arg(0)

	if *outputFile == "" {
		*outputFile = fileNameWithoutExtension(inputFile) + ".vcf"
	}

	return CommandLine{InputFile: inputFile, OutputFile: *outputFile, PrintCSVFields: *printCSVFields}
}

func main() {
	args := parseCommandLine()

	var reader Reader
	var writer Writer
	var err error

	inputFile := args.InputFile
	switch filepath.Ext(inputFile) {
	case ".csv":
		reader, err = NewCSVReader(inputFile)
	case ".txt":
		reader, err = NewTxtReader(inputFile)
	default:
		fmt.Fprintln(os.Stderr, "Unknown file extension. See --help")
		os.Exit(1)
	}
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	writer, err = NewVCFWriter(args.OutputFile)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	for {
		p, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else if err == ErrEmptyLine {
				continue
			}
		}
		writer.Write(p)
	}
}
