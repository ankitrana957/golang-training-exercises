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
		name    string
		fields  fields
		wantErr error
		sub     models.Record

		mockCalls []interface{}
	}{
		{name: "Successfully Inserted", fields: fields{rs: mockdb}, sub: models.Record{Student: "Ankit", RollNo: 1, Subject: "Science", Id: 1}, mockCalls: []interface{}{
			mockdb.EXPECT().InsertRecord(gomock.Any()).Return(nil),
		}},
		{name: "Failed Insertion", fields: fields{rs: mockdb}, sub: models.Record{Student: "Ankit", RollNo: 1, Subject: "Science", Id: 1}, mockCalls: []interface{}{
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
