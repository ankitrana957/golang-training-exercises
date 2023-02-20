package processes

import (
	"reflect"
	"testing"

	"github.com/personhashing/models"
)

func TestProcessMsg(t *testing.T) {

	tests := []struct {
		name string
		m    models.JSON
		want models.JSON
	}{
		{name: "Create Msg Signature", m: models.JSON{"id": 1, "name": "Ankit", "age": 22, "phone": "77d70340ee9fd81057ed460d4ccf0f81bb33e5e4408e976bf3502a18d825f54e"}, want: models.JSON{"value": "1Ankit2277d70340ee9fd81057ed460d4ccf0f81bb33e5e4408e976bf3502a18d825f54e                            "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessMsg(tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}
