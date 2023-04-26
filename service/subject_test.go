package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func TestSubjectGetValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMocksubjectstore(ctrl)
	s := models.Subject{Name: "Science", Id: 1}

	type args struct {
		id int
	}
	tests := []struct {
		name     string
		args     args
		want     models.Subject
		wantErr  error
		mockCall []interface{}
	}{
		{name: "Received Records Successfully", args: args{id: 1}, want: models.Subject{Name: "Science", Id: 1}, mockCall: []interface{}{
			mockdb.EXPECT().GetSubject(1).Return(s, nil),
		}},
		{name: "Didn't Received Records", args: args{id: 1}, wantErr: errors.New("Subject didn't found"), mockCall: []interface{}{
			mockdb.EXPECT().GetSubject(1).Return(models.Subject{}, errors.New("Subject didn't found")),
		}},
		{name: "Wrong Id", args: args{id: -1}, wantErr: errors.New("Id is mandatory")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv := NewSubStore(mockdb)
			got, err := serv.GetValidation(tt.args.id)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SubjectService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubjectService.GetValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubjectInsertValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMocksubjectstore(ctrl)

	tests := []struct {
		name      string
		wantErr   error
		s         models.Subject
		mockCalls []interface{}
	}{
		{
			name: "Successful Insertion", s: models.Subject{Name: "Science", Id: 1}, mockCalls: []interface{}{
				mockdb.EXPECT().InsertSubject(gomock.All()).Return(nil),
			},
		},
		{
			name: "Improper Values", wantErr: errors.New("Id and name are mandatory"), s: models.Subject{Name: "", Id: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv := NewSubStore(mockdb)
			err := serv.InsertValidation(tt.s)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Expected %v , Got %v", tt.wantErr, err)
			}
		})
	}
}
