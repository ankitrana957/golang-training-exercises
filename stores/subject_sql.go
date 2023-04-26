package stores

import (
	"database/sql"

	"github.com/student-api/models"
)

type SubjectStore struct {
	*sql.DB
}

// Get Subject with the given id
func (db SubjectStore) GetSubject(id int) (models.Subject, error) {
	row := db.QueryRow(`SELECT * FROM subject where id=?`, id)
	res := models.Subject{}
	err := row.Scan(&res.Name, &res.Id)
	if err != nil {
		return models.Subject{}, err
	}
	return res, nil
}

// Insert Subject to the record database
func (db SubjectStore) InsertSubject(sub models.Subject) error {
	_, err := db.Exec(`INSERT INTO subject VALUES (?,?)`, sub.Name, sub.Id)
	if err != nil {
		return err
	}
	return nil
}
