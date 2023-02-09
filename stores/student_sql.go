package stores

import (
	"database/sql"
	"fmt"

	models "github.com/student-api/models"
)

type SqlDb struct {
	*sql.DB
}

func (db SqlDb) GetStudent(rollNo string) (models.Student, error) {
	var m models.Student
	query := fmt.Sprintf("SELECT * FROM studentDetails WHERE rollNo = %v", rollNo)
	row := db.QueryRow(query)
	err2 := row.Scan(&m.Name, &m.Age, &m.RollNo)
	if err2 != nil {
		return models.Student{}, err2
	}
	return m, nil
}

func (db SqlDb) InsertStudent(p models.Student) error {
	_, err := db.Exec(`INSERT INTO studentDetails VALUES (?,?,?)`, p.Name, p.Age, p.RollNo)
	if err != nil {
		return err
	}
	return nil
}
