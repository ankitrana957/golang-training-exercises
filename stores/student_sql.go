package stores

import (
	"database/sql"
	"fmt"

	models "github.com/student-api/models"
)

type Sqldb struct {
	*sql.DB
}

func (db Sqldb) GetAll() ([]models.Student, error) {
	var s []models.Student
	rows, err := db.Query("Select * from studentDetails")
	if err != nil {
		return []models.Student{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var m models.Student
		err1 := rows.Scan(&m.Name, &m.Age, &m.RollNo)
		if err1 != nil {
			return []models.Student{}, err1
		}
		s = append(s, m)
	}
	return s, nil
}

func (db Sqldb) Get(rollNo string) (models.Student, error) {
	var m models.Student
	query := fmt.Sprintf("Select * from studentDetails where rollNo = %v", rollNo)
	row := db.QueryRow(query)
	err2 := row.Scan(&m.Name, &m.Age, &m.RollNo)
	if err2 != nil {
		return models.Student{}, err2
	}
	return m, nil
}

func (db Sqldb) Delete(rollNo string) error {
	query := fmt.Sprintf(`Delete from studentDetails where RollNo=%s`, rollNo)
	_, err := db.Exec(query)
	return err
}

func (db Sqldb) Insert(p models.Student) error {
	query := fmt.Sprintf(`INSERT INTO studentDetails VALUES ("%s",%d,%d)`, p.Name, p.Age, p.RollNo)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (db Sqldb) Update(s models.Student) error {
	query := fmt.Sprintf(`UPDATE studentDetails SET name='%s',age=%d WHERE rollNo=%d;`, s.Name, s.Age, s.RollNo)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
