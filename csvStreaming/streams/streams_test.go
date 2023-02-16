package streams

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/personhashing/models"
)

func TestProcessHashing(t *testing.T) {
	p := []models.Person{
		{
			Id:    2,
			Name:  "test1",
			Age:   1,
			Phone: "123456",
		},
	}
	tests := []struct {
		name string
		want []models.Person
	}{
		{name: "HashPhone", want: []models.Person{{Phone: "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", Id: 2,
			Name: "test1",
			Age:  1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(chan models.Person)
			var output []models.Person
			go func() {
				defer close(input)
				for i := range p {
					input <- p[i]
				}
			}()
			got := processHashing(input)
			for i := range got {
				output = append(output, i)
			}
			if !reflect.DeepEqual(output, tt.want) {
				t.Errorf("processHashing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessMsg(t *testing.T) {
	p := []models.Person{
		{
			Phone: "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92",
			Id:    2,
			Name:  "test1",
			Age:   1,
		}}
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Msg Signature Encoding", want: []string{"2test118d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92                             "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make(chan models.Person)
			var output []string
			go func() {
				defer close(input)
				for i := range p {
					input <- p[i]
				}
			}()
			got := processMsg(input)
			for i := range got {
				output = append(output, i)
			}

			if !reflect.DeepEqual(output, tt.want) {
				t.Errorf("processMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPipelining(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{name: "Pipelining", args: args{strings.NewReader("1,Ankit,22,38279332")}, wantOutput: "1Ankit220865affc9d12a2b0760e5aef02a3a063939fcc0bb83c436a50e8775194219eff                            \n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			Pipelining(tt.args.input, output)
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("Pipelining() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
