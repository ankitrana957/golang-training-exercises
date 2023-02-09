package service

import (
	"errors"

	"github.com/student-api/models"
)

type subjectstore interface {
	GetSubject(id int) (models.Subject, error)
	InsertSubject(sub models.Subject) error
}

func NewSubStore(db subjectstore) SubjectService {
	return SubjectService{db}
}

type SubjectService struct {
	db subjectstore
}

func (serv SubjectService) GetValidation(id int) (models.Subject, error) {
	if id > 0 {
		return serv.db.GetSubject(id)
	}
	return models.Subject{}, errors.New("Id is mandatory")
}

func (serv SubjectService) InsertValidation(sub models.Subject) error {
	if sub.Id > 0 || sub.Name != "" {
		return serv.db.InsertSubject(sub)
	}
	return errors.New("Id and name are mandatory")
}
