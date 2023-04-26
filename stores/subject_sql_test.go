package stores

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	models "github.com/student-api/models"
)

func TestGetSubject(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		want     models.Subject
		wantErr  error
		mockCall func(mock sqlmock.Sqlmock)
	}{
		{name: "Get Subject record", id: 1, want: models.Subject{Name: "Science", Id: 1}, mockCall: func(mock sqlmock.Sqlmock) {
			rs := sqlmock.NewRows([]string{"name", "id"}).AddRow("Science", 1)
			mock.ExpectQuery("SELECT *").WillReturnRows(rs)
		}},
		{name: "Multiple parameters found", id: 1, want: models.Subject{Name: "Science", Id: 1}, mockCall: func(mock sqlmock.Sqlmock) {
			rs := sqlmock.NewRows([]string{"name", "age", "id"}).AddRow("Ankit", 22, 8)
			mock.ExpectQuery("SELECT *").WillReturnRows(rs).WillReturnError(errors.New("Multiple parameters found"))
		}, wantErr: errors.New("Multiple parameters found")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			sql_db := SubjectStore{db}
			tt.mockCall(mock)
			got, err := sql_db.GetSubject(tt.id)
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

func TestInsertSubject(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		obj      models.Subject
		wantErr  error
		mockCall func(mock sqlmock.Sqlmock)
	}{
		{
			name: "Insertion Successful", obj: models.Subject{Name: "Science", Id: 8}, mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
			}},
		{
			name: "Insertion Failed", obj: models.Subject{Name: "Science", Id: 8}, wantErr: errors.New("Failed to insert subject"), mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("Failed to insert subject"))
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			sql_db := SubjectStore{db}
			tt.mockCall(mock)
			err := sql_db.InsertSubject(tt.obj)
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("Sqldb.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
