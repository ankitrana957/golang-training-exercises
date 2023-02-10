package stores

import (
	"github.com/student-api/models"
)

// Insert Record to the record database
func (db SqlDb) InsertRecord(sub models.Record) error {
	_, err := db.Exec(`INSERT INTO record VALUES (?,?)`, sub.RollNo, sub.Id)
	if err != nil {
		return err
	}
	return nil
}

// Get all the subject ids enrolled with the student rollNo
func (db SqlDb) GetAllSubjects(rollNo string) ([]int, error) {
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
