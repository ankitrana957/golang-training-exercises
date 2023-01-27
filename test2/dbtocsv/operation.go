package dbtocsv

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Name   string
	Age    int
	RollNo int
}

func openConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	if err != nil {
		return nil, errors.New("Error Connecting to database")
	}
	return db, nil
}

func retriveDataFromDB(db *sql.DB) (s []student) {
	rows, _ := db.Query("SELECT * FROM studentDetails")
	defer rows.Close()
	for rows.Next() {
		var stud student
		var phn1, phn2 string
		rows.Scan(&stud.Name, &stud.Age, &stud.RollNo, &phn1, &phn2)
		s = append(s, stud)
	}
	return s
}

func createFile(s string) *os.File {
	f, _ := os.Create(s)
	return f
}

func writeFile(s student, w io.Writer) error {
	a := fmt.Sprintf("%d, %s, %d \n", s.RollNo, s.Name, s.Age)
	_, err := w.Write([]byte(a))
	if err != nil {
		return errors.New("Error in writing buffer ")
	}
	return nil
}

func process(s []student, f *os.File) {
	for _, c := range s {
		writeFile(c, f)
	}
}

// func main() {
// 	db, _ := openConnection()
// 	data := retriveDataFromDB(db)
// 	f := createFile("data.txt")
// 	process(data, f)

// }
