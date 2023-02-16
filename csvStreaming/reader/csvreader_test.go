package reader

import (
	"encoding/csv"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/personhashing/models"
)

type InvalidReader string

func (r InvalidReader) Read(bytes []byte) (int, error) {
	return 0, errors.New("Invalid Reader")
}

func TestCsvReaderRead(t *testing.T) {
	type fields struct {
		Reader *csv.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.Person
		wantErr error
	}{
		{name: "Retriving Person Successful", fields: fields{Reader: csv.NewReader(strings.NewReader(`1,Ankit,22,8700917756`))}, want: models.Person{Name: "Ankit", Id: 1, Age: 22, Phone: "8700917756"}},
		{name: "Invalid Id", fields: fields{Reader: csv.NewReader(strings.NewReader(`a,Ankit,22,8700917756`))}, wantErr: errors.New("Invalid data type conversion")},
		{name: "Invalid Age", fields: fields{Reader: csv.NewReader(strings.NewReader(`1,Ankit,a,8700917756`))}, wantErr: errors.New("Invalid data type conversion")},
		{name: "Invalid Reader", fields: fields{Reader: csv.NewReader(InvalidReader(`1Ankita8700917756,,tfhf`))}, wantErr: errors.New("Invalid Reader")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CsvReader{
				Reader: tt.fields.Reader,
			}
			got, err := c.Read()
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("CsvReader.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CsvReader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
