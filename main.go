package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
)

const (
	coverageDir = "cov/"
)

type procInfo struct {
	args []byte
	pid  int
}

var inputFile string
var targetProg string
var covAddrs map[string]bool

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	flag.StringVar(&inputFile, "input", "", "File containing initial corpus.")
	flag.StringVar(&targetProg, "target", "", "Target program to be fuzzed.")
}

func main() {
	flag.Parse()

	if inputFile == "" || targetProg == "" {
		fmt.Printf("Usage: %s -input {inputFile} -target {targetProgram}\n", os.Args[0])
		return
	}

	covAddrs = make(map[string]bool)
	sigLog := make(chan procInfo)
	sigExec := make(chan procInfo)
	sigWork := make(chan []byte)

	os.MkdirAll(coverageDir+"tmp", os.ModePerm)

	inFile, err := ioutil.ReadFile(inputFile)
	checkErr(err)
	corpus := make([]byte, len(inFile))
	copy(corpus, inFile)

	fmt.Printf("CORPUS: % 0x\n", corpus)
	go logger(sigLog, sigExec)
	go worker(sigLog, sigExec, sigWork)
	go generator(corpus, sigWork)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	<-c
	os.RemoveAll(coverageDir + "tmp")
}
