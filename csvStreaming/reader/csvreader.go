package reader

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"

	"github.com/personhashing/models"
)

type csvReader struct {
	*csv.Reader
}

func NewCSVReader(r io.Reader) *csvReader {
	return &csvReader{csv.NewReader(r)}
}

func (c csvReader) Read() (models.JSON, error) {
	record, err := c.Reader.Read()
	if err != nil {
		return models.JSON{}, err
	}

	id, err1 := strconv.Atoi(record[0])
	if err1 != nil {
		return models.JSON{}, errors.New("Invalid data type conversion")
	}

	age, err2 := strconv.Atoi(record[2])
	if err2 != nil {
		return models.JSON{}, errors.New("Invalid data type conversion")
	}

	// Return a Json Object
	jsonObj := models.JSON{"id": id, "name": record[1], "age": age, "phone": record[3]}

	return jsonObj, nil
}
