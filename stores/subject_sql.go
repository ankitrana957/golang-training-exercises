package stores

import (
	"github.com/student-api/models"
)

// Get Subject with the given id
func (db SqlDb) GetSubject(id int) (models.Subject, error) {
	row := db.QueryRow(`SELECT * FROM subject where id=?`, id)
	res := models.Subject{}
	err := row.Scan(&res.Name, &res.Id)
	if err != nil {
		return models.Subject{}, err
	}
	return res, nil
}

// Insert Subject to the record database
func (db SqlDb) InsertSubject(sub models.Subject) error {
	_, err := db.Exec(`INSERT INTO subject VALUES (?,?)`, sub.Name, sub.Id)
	if err != nil {
		return err
	}
	return nil
}
