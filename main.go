package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inFile = kingpin.Arg("file", "TOML file.").Required().ExistingFile()
	quiet  = kingpin.Flag("quiet", "Don't output on success.").Short('q').Bool()
)

func main() {

	// support -h for --help
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	data, _ := ioutil.ReadFile(*inFile)

	var f interface{}
	_, err := toml.Decode(string(data), &f)
	if err != nil {
		fmt.Println("ERROR:", *inFile, err)
		os.Exit(1)
	}

	if !*quiet {
		fmt.Println("OK:", *inFile)
	}
}
