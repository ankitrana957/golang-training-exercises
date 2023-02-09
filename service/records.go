package service

import (
	"github.com/student-api/models"
)

type recordstore interface {
	InsertRecord(sub models.Record) error
}

type enrollmentService struct {
	rs recordstore
}

// Factory
func NewEnrollmentStore(r recordstore) enrollmentService {
	return enrollmentService{r}
}

func (e enrollmentService) Insert(sub models.Record) error {
	err := e.rs.InsertRecord(sub)
	if err != nil {
		return err
	}
	return nil
}
