package writer

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/personhashing/models"
)

type InvalidWriter string

func (w InvalidWriter) Write(bytes []byte) (int, error) {
	return 0, errors.New("Invalid Writer found")
}

func TestCsvWriter(t *testing.T) {
	tests := []struct {
		name    string
		writer  io.Writer
		wantErr error
		data    models.JSON
	}{
		{
			name: "Writing Data", writer: &bytes.Buffer{}, data: models.JSON{"value": "1Ankit220865affc9d12a2b0760e5aef02a3a063939fcc0bb83c436a50e8775194219eff                            \n"},
		},
		{
			name: "Invalid Writer", writer: InvalidWriter("Some Data"), wantErr: errors.New("Invalid writer found"), data: models.JSON{"value": ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewCSVWriter(tt.writer)
			err := w.Write(tt.data)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("csvWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
