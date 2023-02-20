package main

import (
	"log"
	"os"

	"github.com/personhashing/models"
	"github.com/personhashing/pipelining"
	"github.com/personhashing/processes"
	"github.com/personhashing/reader"

	"github.com/personhashing/writer"
)

func main() {

	f, err1 := os.Open("input.csv")

	if err1 != nil {
		log.Fatal(err1)
	}
	defer f.Close()

	outputFile, _ := os.Create("output.csv")
	defer outputFile.Close()

	r := reader.NewCSVReader(f)
	w := writer.NewCSVWriter(outputFile)
	a := []models.Process{processes.ProcessHashing, processes.ProcessMsg}
	err := pipelining.PipelineProcessor(r, w, a)

	if err != nil {
		log.Fatal(err)
	}
}
