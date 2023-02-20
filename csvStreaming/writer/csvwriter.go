package writer

import (
	"encoding/csv"
	"io"

	"github.com/personhashing/models"
)

type csvWriter struct {
	*csv.Writer
}

func NewCSVWriter(w io.Writer) *csvWriter {
	return &csvWriter{csv.NewWriter(w)}
}

func (w csvWriter) Write(data models.JSON) error {
	a := []string{data["value"].(string)}
	w.Writer.Write(a)
	defer w.Flush()
	return nil
}
