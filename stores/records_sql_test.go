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
		{name: "Successful insertion of record", obj: models.Record{RollNo: 1, Id: 2}, mockCall: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
		}, wantErr: nil},
		{name: "Unsuccessful insertion of record", obj: models.Record{RollNo: 1, Id: 2}, mockCall: func(mock sqlmock.Sqlmock) {
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

func TestGetAllSubs(t *testing.T) {
	tests := []struct {
		name     string
		rollNo   string
		wantErr  error
		mockCall func(mock sqlmock.Sqlmock)
		want     []int
	}{
		{name: "Got Records", rollNo: "1", mockCall: func(mock sqlmock.Sqlmock) {
			rs := mock.NewRows([]string{"id"}).AddRow(1).AddRow(2)
			mock.ExpectQuery("SELECT id").WillReturnRows(rs)
		}, want: []int{1, 2}},
		{name: "Didn't get any records", rollNo: "1", mockCall: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery("SELECT id").WillReturnError(errors.New("Student is not enrolled in any subject"))
		}, wantErr: errors.New("Student is not enrolled in any subject")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			sql_db := SqlDb{db}
			tt.mockCall(mock)
			got, err := sql_db.GetAllSubjects(tt.rollNo)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Expected %s, Got %s", tt.wantErr.Error(), err.Error())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, Got %v", tt.want, got)
			}
		})

	}
}
