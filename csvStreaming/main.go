package main

import (
	"fmt"
	"os"

	"github.com/personhashing/filehandling"
	"github.com/personhashing/hashing"
	"github.com/personhashing/models"
	"github.com/personhashing/msgpadding"
)

func main() {
	f, err1 := os.Open("input.csv")
	outputFile, err2 := os.Create("output.csv")

	if err1 != nil && err2 != nil {
		fmt.Print(err1)
	}

	c := make(chan models.Person)
	m := make(chan models.Person)
	k := make(chan string)

	go filehandling.ReadFromCSV(f, c)
	go hashing.HashString(c, m)
	go msgpadding.MakeMsg(m, k)

	for i := range k {
		filehandling.WriteToCSV(outputFile, i)
	}
}
