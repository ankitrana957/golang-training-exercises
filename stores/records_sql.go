package stores

import (
	"github.com/student-api/models"
)

func (db SqlDb) InsertRecord(sub models.Record) error {
	_, err := db.Exec(`INSERT INTO record VALUES (?,?,?,?)`, sub.Student, sub.RollNo, sub.Subject, sub.Id)
	if err != nil {
		return err
	}
	return nil
}
