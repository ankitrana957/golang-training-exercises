package streams

import (
	"crypto/sha256"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/personhashing/models"
	"github.com/personhashing/msgdigest"
	"github.com/personhashing/reader"
	"github.com/personhashing/readwrite"
)

func processHashing(c <-chan models.Person) <-chan models.Person {
	m := make(chan models.Person)
	go func() {
		defer close(m)
		for i := range c {
			phoneHash := fmt.Sprintf("%x", sha256.Sum256([]byte(i.Phone)))
			a := i.SetHashPhone(phoneHash)
			m <- a
		}
	}()
	return m
}

func processMsg(m <-chan models.Person) <-chan string {
	k := make(chan string)
	go func() {
		defer close(k)
		for j := range m {
			a := msgdigest.CreateMsgSignature(j)
			k <- a
		}
	}()
	return k
}

func Pipelining(input io.Reader, output io.Writer) {
	x := csv.NewReader(input)
	c := readwrite.ReadPerson(reader.CsvReader{x})
	m := processHashing(c)
	k := processMsg(m)

	for i := range k {
		readwrite.WritePerson(output, i)
	}
}
