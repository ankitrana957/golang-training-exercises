package msgdigest

import (
	"testing"

	"github.com/personhashing/models"
)

func TestCreateMsgSignature(t *testing.T) {
	tests := []struct {
		name   string
		person models.Person
		want   string
	}{
		{
			name:   "Create Msg Signature",
			person: models.Person{Id: 1, Name: "Ankit", Age: 22, Phone: "8700917756"},
			want:   "1Ankit228700917756                                                                                  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMsgSignature(tt.person); got != tt.want {
				t.Errorf("CreateMsgSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}
