package stores

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	models "github.com/student-api/models"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name     string
		rollNo   string
		want     models.Student
		wantErr  error
		mockCall func(mock sqlmock.Sqlmock)
	}{
		{name: "Get record", rollNo: "8", want: models.Student{Name: "Ankit", Age: 22, RollNo: 8}, mockCall: func(mock sqlmock.Sqlmock) {
			rs := sqlmock.NewRows([]string{"name", "age", "rollNo"}).AddRow("Ankit", 22, 8)
			mock.ExpectQuery("SELECT *").WillReturnRows(rs)
		}},
		{name: "Multiple parameters found", rollNo: "8", want: models.Student{Name: "Ankit", Age: 22, RollNo: 8}, mockCall: func(mock sqlmock.Sqlmock) {
			rs := sqlmock.NewRows([]string{"name", "age", "rollNo", "phn"}).AddRow("Ankit", 22, 8, "27638721")
			mock.ExpectQuery("SELECT *").WillReturnRows(rs).WillReturnError(errors.New("Multiple parameters found"))
		}, wantErr: errors.New("Multiple parameters found")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			sql_db := SqlDb{db}
			tt.mockCall(mock)
			got, err := sql_db.GetStudent(tt.rollNo)
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("Sqldb.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (tt.wantErr == nil) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sqldb.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		obj      models.Student
		wantErr  error
		mockCall func(mock sqlmock.Sqlmock)
	}{
		{
			name: "Insertion Successful", obj: models.Student{Name: "Ankit", Age: 22, RollNo: 8}, mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
			}},
		{
			name: "Insertion Failed", obj: models.Student{Name: "Ankit", Age: 22, RollNo: 8}, wantErr: errors.New("Failed to insert student"), mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("Failed to insert student"))
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			sql_db := SqlDb{db}
			tt.mockCall(mock)
			err := sql_db.InsertStudent(tt.obj)
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("Sqldb.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
