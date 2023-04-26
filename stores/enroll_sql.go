package stores

import (
	"database/sql"

	"github.com/student-api/models"
)

type EnrollmentStore struct {
	*sql.DB
}

// Insert Record to the record database
func (db EnrollmentStore) InsertRecord(sub models.Enroll) error {
	_, err := db.Exec(`INSERT INTO record VALUES (?,?)`, sub.RollNo, sub.Id)
	if err != nil {
		return err
	}
	return nil
}

// Get all the subject ids enrolled with the student rollNo
func (db EnrollmentStore) GetAllSubjects(rollNo string) ([]int, error) {
	rows, err := db.Query(`SELECT id FROM record WHERE rollNo = ?`, rollNo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var s []int
	for rows.Next() {
		var r int
		rows.Scan(&r)
		s = append(s, r)
	}
	return s, nil
}
