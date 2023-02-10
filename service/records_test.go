package service

import (
	"errors"
	reflect "reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func Test_enrollmentService_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMockrecordstore(ctrl)
	type fields struct {
		rs recordstore
	}
	tests := []struct {
		name      string
		fields    fields
		wantErr   error
		sub       models.Record
		mockCalls []interface{}
	}{
		{name: "Successfully Inserted", fields: fields{rs: mockdb}, sub: models.Record{RollNo: 1, Id: 1}, mockCalls: []interface{}{
			mockdb.EXPECT().InsertRecord(gomock.Any()).Return(nil),
		}},
		{name: "Failed Insertion", fields: fields{rs: mockdb}, sub: models.Record{RollNo: 1, Id: 1}, mockCalls: []interface{}{
			mockdb.EXPECT().InsertRecord(gomock.Any()).Return(errors.New("Error in insertion of record")),
		}, wantErr: errors.New("Error in insertion of record")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			e := enrollmentService{
				rs: tt.fields.rs,
			}
			err := e.Insert(tt.sub)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StudentEnrollmentService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_enrollmentService_GetSubs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMockrecordstore(ctrl)

	tests := []struct {
		name      string
		rollNo    string
		wantErr   error
		want      []int
		rs        recordstore
		mockCalls []interface{}
	}{
		{name: "Get Subs", rs: mockdb, rollNo: "1", want: []int{1, 2, 3}, mockCalls: []interface{}{
			mockdb.EXPECT().GetAllSubjects(gomock.Any()).Return([]int{1, 2, 3}, nil),
		}},
		{name: "Failed to get Subs", rs: mockdb, rollNo: "1", wantErr: errors.New("Fail to get data"), mockCalls: []interface{}{
			mockdb.EXPECT().GetAllSubjects(gomock.Any()).Return(nil, errors.New("Fail to get data")),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := enrollmentService{
				rs: tt.rs,
			}
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
