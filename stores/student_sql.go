package stores

import (
	"database/sql"

	models "github.com/student-api/models"
)

type StudentStore struct {
	*sql.DB
}

// Get Student with the given roll no
func (db StudentStore) GetStudent(rollNo string) (models.Student, error) {
	var m models.Student
	row := db.QueryRow(`SELECT * FROM studentDetails WHERE rollNo = ?`, rollNo)
	err2 := row.Scan(&m.Name, &m.Age, &m.RollNo)
	if err2 != nil {
		return models.Student{}, err2
	}
	return m, nil
}

// Insert Student to the studentDetails database
func (db StudentStore) InsertStudent(p models.Student) error {
	_, err := db.Exec(`INSERT INTO studentDetails VALUES (?,?,?)`, p.Name, p.Age, p.RollNo)
	if err != nil {
		return err
	}
	return nil
}
