package service

import (
	"errors"
	"strconv"

	"github.com/student-api/models"
)

type studentdatastore interface {
	InsertStudent(models.Student) error
	GetStudent(string) (models.Student, error)
}

type StudentEnrollmentService struct {
	db         studentdatastore
	enrollment enrollmentServiceSample
	subject    subjectServiceSample
}

type enrollmentServiceSample interface {
	Insert(sub models.Record) error
}

type subjectServiceSample interface {
	GetValidation(id int) (models.Subject, error)
	InsertValidation(sub models.Subject) error
}

func NewStudentService(db studentdatastore, enroll enrollmentServiceSample, sub subjectServiceSample) StudentEnrollmentService {
	return StudentEnrollmentService{db, enroll, sub}
}

func (s StudentEnrollmentService) GetValidation(rollNo string) (models.Student, error) {
	if rollNo != "" {
		return s.db.GetStudent(rollNo)
	}
	return models.Student{}, errors.New("RollNo is not given")
}

func (s StudentEnrollmentService) PostValidation(data models.Student) error {
	if data.Name != "" && data.RollNo != 0 {
		return s.db.InsertStudent(data)
	}
	return errors.New("RollNo and Name are mandatory")
}

func (s StudentEnrollmentService) Enroll(id, rollNo int) error {
	roll := strconv.Itoa(rollNo)
	stu, err := s.db.GetStudent(roll)
	if err != nil {
		return err
	}
	sub, err1 := s.subject.GetValidation(id)
	if err1 != nil {
		return err1
	}
	record := models.Record{
		Student: stu.Name,
		RollNo:  stu.RollNo,
		Subject: sub.Name,
		Id:      sub.Id,
	}
	err3 := s.enrollment.Insert(record)

	if err3 != nil {
		return err3
	}

	return nil

}
