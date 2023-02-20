package processes

import (
	"reflect"
	"testing"

	"github.com/personhashing/models"
)

func TestProcessHashing(t *testing.T) {
	tests := []struct {
		name string
		want models.JSON
		c    models.JSON
	}{
		{name: "Get Phone Hash", c: models.JSON{"id": 1, "name": "Ankit", "age": 22, "phone": "878515665"}, want: models.JSON{"id": 1, "name": "Ankit", "age": 22, "phone": "77d70340ee9fd81057ed460d4ccf0f81bb33e5e4408e976bf3502a18d825f54e"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessHashing(tt.c)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessHashing() = %v, want %v", got, tt.want)
			}
		})
	}
}
