package student

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	models "github.com/student-api/models"
)

func TestGetAll(t *testing.T) {
	tests := []struct {
		name     string
		want     []models.Student
		mockCall func(mock sqlmock.Sqlmock)
		wantErr  error
	}{
		{name: "Successfull get request", want: []models.Student{
			{Name: "Amit", Age: 21, RollNo: 7}, {Name: "Ankit", Age: 22, RollNo: 9},
		},
			mockCall: func(mock sqlmock.Sqlmock) {
				rs := sqlmock.NewRows([]string{"name", "age", "rollNo"}).AddRow("Amit", 21, 7).AddRow("Ankit", 22, 9)
				mock.ExpectQuery("SELECT *").WillReturnRows(rs)
			},
		},
		{name: "Mismatched query found", want: []models.Student{}, wantErr: errors.New("Mismatched query found"),
			mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT *").WillReturnError(errors.New("Mismatched query found"))
			},
		},
		{name: "Multiple parameters found", want: nil,
			wantErr: errors.New("Multiple arguments found"),

			mockCall: func(mock sqlmock.Sqlmock) {
				rs := sqlmock.NewRows([]string{"name", "age", "rollNo", "phn"}).AddRow("Amit", 21, 7, "26531723576")
				mock.ExpectQuery("SELECT *").WillReturnRows(rs)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			sql_db := Sqldb{db}

			tt.mockCall(mock)
			got, err := sql_db.GetAll()
			if (tt.wantErr != nil) && tt.wantErr == err {
				t.Errorf("Expected %s , got %s", tt.wantErr.Error(), err.Error())
			}
			if (tt.wantErr == nil) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sqldb.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}

}

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
			sql_db := Sqldb{db}
			tt.mockCall(mock)
			got, err := sql_db.Get(tt.rollNo)
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

func TestDelete(t *testing.T) {
	tests := []struct {
		name     string
		rollNo   string
		wantErr  error
		mockCall func(mock sqlmock.Sqlmock)
	}{
		{
			name: "Deletion Successful", rollNo: "8", mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM").WillReturnResult(sqlmock.NewResult(1, 1))
			}},
		{
			name: "Failed Deletion", wantErr: errors.New("Failed to delete student"), mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("DELETE FROM").WillReturnError(errors.New("Failed to delete student"))
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			sqldb := Sqldb{db}
			tt.mockCall(mock)
			err := sqldb.Delete(tt.rollNo)
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("Sqldb.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name     string
		want     models.Student
		wantErr  error
		mockCall func(mock sqlmock.Sqlmock)
	}{
		{
			name: "Insertion Successful", want: models.Student{Name: "Ankit", Age: 22, RollNo: 8}, mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
			}},
		{
			name: "Insertion Failed", want: models.Student{Name: "Ankit", Age: 22, RollNo: 8}, wantErr: errors.New("Failed to insert student"), mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO").WillReturnError(errors.New("Failed to insert student"))
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			sqldb := Sqldb{db}
			tt.mockCall(mock)
			err := sqldb.Insert(tt.want)
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("Sqldb.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name     string
		want     models.Student
		wantErr  error
		mockCall func(mock sqlmock.Sqlmock)
	}{
		{
			name: "Updated", want: models.Student{Name: "Ankit", Age: 22, RollNo: 8}, mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE studentDetails").WillReturnResult(sqlmock.NewResult(0, 1))
			}},
		{
			name: "Update Failed", want: models.Student{Name: "Ankit", Age: 22, RollNo: 8}, wantErr: errors.New("Updation Failed"), mockCall: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE studentDetails").WillReturnError(errors.New("Updation Failed"))
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			sqldb := Sqldb{db}
			tt.mockCall(mock)
			err := sqldb.Update(tt.want)
			if (err != nil) && err.Error() != tt.wantErr.Error() {
				t.Errorf("Sqldb.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
