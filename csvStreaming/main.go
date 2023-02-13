package main

import (
	"log"
	"os"

	"github.com/personhashing/filehandling"
	"github.com/personhashing/hashing"
	"github.com/personhashing/models"
	"github.com/personhashing/msgpadding"
)

func ServeHashing(c, m chan models.Person) {
	for i := range c {
		a := hashing.HashString(i)
		m <- a
	}
	close(m)
}

func ServeMsg(m chan models.Person, k chan string) {
	for j := range m {
		a := msgpadding.MakeMsg(j)
		k <- a
	}
	close(k)
}

func main() {
	f, err1 := os.Open("input.csv")

	if err1 != nil {
		log.Fatal(err1)
	}
	defer f.Close()

	c := make(chan models.Person)
	m := make(chan models.Person)
	k := make(chan string)

	outputFile, _ := os.Create("output.csv")
	defer outputFile.Close()

	go filehandling.ReadFromCSV(f, c)
	go ServeHashing(c, m)
	go ServeMsg(m, k)

	for i := range k {
		filehandling.WriteToCSV(outputFile, i)
	}
}
