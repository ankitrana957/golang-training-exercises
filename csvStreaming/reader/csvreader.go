package reader

import (
	"encoding/csv"
	"errors"
	"strconv"

	"github.com/personhashing/models"
)

type CsvReader struct {
	*csv.Reader
}

func (c CsvReader) Read() (models.Person, error) {
	record, err := c.Reader.Read()
	if err != nil {
		return models.Person{}, err
	}

	id, err1 := strconv.Atoi(record[0])
	if err1 != nil {
		return models.Person{}, errors.New("Invalid data type conversion")
	}

	age, err2 := strconv.Atoi(record[2])
	if err2 != nil {
		return models.Person{}, errors.New("Invalid data type conversion")
	}

	return models.Person{Id: id, Name: record[1], Age: age, Phone: record[3]}, nil
}
