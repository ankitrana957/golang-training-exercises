package worker

import (
	"reflect"
	"testing"

	"github.com/personhashing/models"
)

func TestWorker(t *testing.T) {

	tests := []struct {
		args  []models.JSON
		input models.Process
		want  models.JSON
	}{
		{
			args: []models.JSON{{"Phone": "234"}},
			input: models.Process(func(m models.JSON) models.JSON {
				return models.JSON{"Phone": "123"}
			}),
			want: models.JSON{"Phone": "123"},
		},
	}

	for _, tt := range tests {

		InCh := make(chan models.JSON)
		OutCh := make(<-chan models.JSON)
		// temp := make(chan models.JSON)

		OutCh = Worker(tt.input, InCh)

		for _, value := range tt.args {

			InCh <- value
		}
		close(InCh)

		for value := range OutCh {

			if !reflect.DeepEqual(value, tt.want) {
				t.Errorf("got %v, want %v", value, tt.want)
			}
		}

	}
}
