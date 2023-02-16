package readwrite

import (
	"io"

	"github.com/personhashing/models"
)

type PersonReader interface {
	Read() (models.Person, error)
}

func ReadPerson(data PersonReader) <-chan models.Person {
	c := make(chan models.Person)
	go func() {
		defer close(c)
		for {
			record, err := data.Read()
			if err != nil {
				return
			}
			c <- record
		}
	}()
	return c
}

func WritePerson(f io.Writer, data string) (int, error) {
	n, err := f.Write([]byte(data + "\n"))
	if err != nil {
		return 0, err
	}
	return n, nil
}
