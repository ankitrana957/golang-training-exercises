package filehandling

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"

	"github.com/personhashing/models"
)

func ReadFromCSV(data io.Reader, c chan models.Person) {
	r := csv.NewReader(data)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			if errors.Is(err, csv.ErrFieldCount) {
				log.Fatal("wrong fields")
			} else {
				log.Fatal(err)
			}
		}
		id, err1 := strconv.Atoi(record[0])
		age, err2 := strconv.Atoi(record[2])
		if err1 != nil && err2 != nil {
			continue
		}
		p := models.Person{Id: id, Name: record[1], Age: age, Phone: record[3]}
		c <- p
	}
	close(c)
}

func WriteToCSV(f io.Writer, data string) {
	f.Write([]byte(data + "\n"))
}
