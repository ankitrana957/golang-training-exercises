package stores

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/student-api/models"
)

func TestInsertRecord(t *testing.T) {
	tests := []struct {
		name     string
		obj      models.Record
		wantErr  error
		mockCall func(mock sqlmock.Sqlmock)
	}{
		{name: "Successful insertion of record", obj: models.Record{Student: "Ankit", RollNo: 1, Subject: "Science", Id: 2}, mockCall: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
		}, wantErr: nil},
		{name: "Unsuccessful insertion of record", obj: models.Record{Student: "Ankit", RollNo: 1, Subject: "Science", Id: 2}, mockCall: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec("INSERT").WillReturnError(errors.New("Failed Insertion"))
		}, wantErr: errors.New("Failed Insertion")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			sql_db := SqlDb{db}
			tt.mockCall(mock)
			err := sql_db.InsertRecord(tt.obj)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Expected %s, Got %s", tt.wantErr.Error(), err.Error())
			}
		})

	}
}
