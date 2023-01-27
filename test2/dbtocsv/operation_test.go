package dbtocsv

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestRetriveDataFromDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error("Error connecting to database")
	}
	tests := []struct {
		name  string
		db    *sql.DB
		wantS []student
	}{
		{name: "Successful Compilation", db: db, wantS: []student{
			{Name: "Ankit", Age: 20, RollNo: 1},
		}},
	}

	rs := mock.NewRows([]string{"name", "age", "rollNo"}).AddRow("Ankit", 20, 1)
	mock.ExpectQuery("SELECT *").WillReturnRows(rs)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := retriveDataFromDB(tt.db); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("retriveDataFromDB() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
