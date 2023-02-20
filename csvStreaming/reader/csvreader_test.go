package reader

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/personhashing/models"
)

type InvalidReader string

func (r InvalidReader) Read([]byte) (int, error) {
	return 0, errors.New("Invalid reader found")
}

func TestCsvReader_Read(t *testing.T) {
	tests := []struct {
		name    string
		reader  io.Reader
		want    models.JSON
		wantErr error
	}{
		{name: "Succesfully Data Read", reader: strings.NewReader(`1,Ankit,22,878515665`), want: models.JSON{"id": 1, "name": "Ankit", "age": 22, "phone": "878515665"}, wantErr: nil},
		{name: "Invalid Id", reader: strings.NewReader(`a,Ankit,22,8700917756`), wantErr: errors.New("Invalid data type conversion"), want: models.JSON{}},
		{name: "Invalid Age", reader: strings.NewReader(`1,Ankit,a,8700917756`), wantErr: errors.New("Invalid data type conversion"), want: models.JSON{}},
		{name: "Invalid Reader", reader: InvalidReader("Some Data"), wantErr: errors.New("Invalid reader found"), want: models.JSON{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCSVReader(tt.reader)
			got, err := c.Read()
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("csvReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("csvReader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
