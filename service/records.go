package service

import (
	"github.com/student-api/models"
)

type recordstore interface {
	InsertRecord(sub models.Record) error
	GetAllSubjects(rollNo string) ([]int, error)
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

func (e enrollmentService) GetSubs(rollNo string) ([]int, error) {
	a, err := e.rs.GetAllSubjects(rollNo)
	if err != nil {
		return nil, err
	}
	return a, nil
}
