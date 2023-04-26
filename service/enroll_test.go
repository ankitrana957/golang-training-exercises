package service

import (
	"errors"
	reflect "reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func TestRecordInsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMockrecordstore(ctrl)
	tests := []struct {
		name      string
		wantErr   error
		sub       models.Enroll
		mockCalls []interface{}
	}{
		{name: "Successfully Inserted", sub: models.Enroll{RollNo: 1, Id: 1}, mockCalls: []interface{}{
			mockdb.EXPECT().InsertRecord(gomock.Any()).Return(nil),
		}},
		{name: "Failed Insertion", sub: models.Enroll{RollNo: 1, Id: 1}, mockCalls: []interface{}{
			mockdb.EXPECT().InsertRecord(gomock.Any()).Return(errors.New("Error in insertion of record")),
		}, wantErr: errors.New("Error in insertion of record")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEnrollmentStore(mockdb)
			err := e.Insert(tt.sub)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StudentEnrollmentService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRecordGetSubs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMockrecordstore(ctrl)

	tests := []struct {
		name      string
		rollNo    string
		wantErr   error
		want      []int
		mockCalls []interface{}
	}{
		{name: "Get Subs", rollNo: "1", want: []int{1, 2, 3}, mockCalls: []interface{}{
			mockdb.EXPECT().GetAllSubjects(gomock.Any()).Return([]int{1, 2, 3}, nil),
		}},
		{name: "Failed to get Subs", rollNo: "1", wantErr: errors.New("Fail to get data"), mockCalls: []interface{}{
			mockdb.EXPECT().GetAllSubjects(gomock.Any()).Return(nil, errors.New("Fail to get data")),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEnrollmentStore(mockdb)
			got, err := e.GetSubs(tt.rollNo)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("enrollmentService.GetSubs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("enrollmentService.GetSubs() = %v, want %v", got, tt.want)
			}
		})
	}
}
