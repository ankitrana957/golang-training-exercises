package writer

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/personhashing/models"
)

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
