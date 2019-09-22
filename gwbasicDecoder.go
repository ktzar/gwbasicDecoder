package main

import (
    "fmt"
    "gwbasicParser"
    "io/ioutil"
    "flag"
    "os"
)

func main() {
    var inFile string
    var lineNumbers bool

    flag.StringVar(&inFile, "in", "", "File to parse (required)")
    flag.BoolVar(&lineNumbers, "linenumbers", false, "Show line numbers")
	flag.Parse()

    if inFile == "" {
		fmt.Println("Missing input file")
		flag.PrintDefaults()
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(inFile)
	if err != nil {
		fmt.Println("Can't read the file", inFile)
		os.Exit(1)
	}

    program, err := gwbasicParser.ParseProgram(data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
    if lineNumbers {
        fmt.Print(program.WithLines())
    } else {
        fmt.Print(program)
    }
}
