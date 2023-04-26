package service

import (
	"github.com/student-api/models"
)

type recordstore interface {
	InsertRecord(sub models.Enroll) error
	GetAllSubjects(rollNo string) ([]int, error)
}

type enrollmentService struct {
	rs recordstore
}

// Factory
func NewEnrollmentStore(r recordstore) enrollmentService {
	return enrollmentService{r}
}

func (e enrollmentService) Insert(sub models.Enroll) error {
	return e.rs.InsertRecord(sub)

}

func (e enrollmentService) GetSubs(rollNo string) ([]int, error) {
	return e.rs.GetAllSubjects(rollNo)
}
