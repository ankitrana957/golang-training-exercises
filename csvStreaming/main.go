package main

import (
	"log"
	"os"

	"github.com/personhashing/streams"
)

func main() {
	f, err1 := os.Open("input.csv")

	if err1 != nil {
		log.Fatal(err1)
	}
	defer f.Close()

	outputFile, _ := os.Create("output.csv")
	defer outputFile.Close()

	// Process Streams
	streams.Pipelining(f, outputFile)
}
