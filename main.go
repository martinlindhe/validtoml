package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	termutil "github.com/andrew-d/go-termutil"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inFile = kingpin.Arg("file", "TOML file.").String()
	quiet  = kingpin.Flag("quiet", "Don't output on success.").Short('q').Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	data, err := readPipeOrFile(*inFile)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	filename := "-"
	if *inFile != "" {
		filename = *inFile
	}

	var f interface{}
	_, err = toml.Decode(string(data), &f)
	if err != nil {
		fmt.Println("ERROR:", filename, err)
		os.Exit(1)
	}

	if !*quiet {
		fmt.Println("OK:", filename)
	}
}

// readPipeOrFile reads from stdin if pipe exists, else from provided file
func readPipeOrFile(fileName string) ([]byte, error) {
	if !termutil.Isatty(os.Stdin.Fd()) {
		return ioutil.ReadAll(os.Stdin)
	}
	if fileName == "" {
		return nil, fmt.Errorf("no piped data and no file provided")
	}
	return ioutil.ReadFile(fileName)
}
