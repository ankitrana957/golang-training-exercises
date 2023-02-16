package readwrite

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/personhashing/models"
)

func TestReadPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := NewMockPersonReader(ctrl)
	type args struct {
		data PersonReader
	}
	tests := []struct {
		name     string
		args     args
		want     []models.Person
		mockCall []interface{}
	}{
		{
			name: "Base Case",
			args: args{data: mockReader},
			mockCall: []interface{}{
				mockReader.EXPECT().Read().Return(models.Person{Name: "Ankit", Age: 22, Id: 1, Phone: "8700917756"}, nil),
				mockReader.EXPECT().Read().Return(models.Person{}, errors.New("End Of File")),
			},
			want: []models.Person{
				{
					Name: "Ankit", Age: 22, Id: 1, Phone: "8700917756"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReadPerson(tt.args.data)
			var output []models.Person
			for i := range got {
				output = append(output, i)
			}
			if !reflect.DeepEqual(output, tt.want) {
				t.Errorf("ReadPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

type InvalidWriter string

func (w InvalidWriter) Write([]byte) (n int, err error) {
	return 0, errors.New("Failed to write data")
}

func TestWritePerson(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		f       io.Writer
		wantF   string
		wantErr error
	}{
		{name: "Testing Writer", args: args{data: "Hello There"}, want: 12, wantF: "Hello There\n", f: &bytes.Buffer{}},
		{name: "Failed to write", wantErr: errors.New("Failed to write data"), args: args{data: "Some Data"}, f: InvalidWriter("wdhjawd")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WritePerson(tt.f, tt.args.data)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("WritePerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WritePerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
