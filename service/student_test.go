package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func TestGetValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockdb := NewMockstudentdatastore(ctrl)
	type args struct {
		rollNo string
	}
	tests := []struct {
		name      string
		args      args
		want      models.Student
		mockCalls []interface{}
		wantErr   error
	}{
		{name: "Success", args: args{rollNo: "1"}, want: models.Student{Name: "Ankit", Age: 21, RollNo: 3}, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent("1").Return(models.Student{Name: "Ankit", Age: 21, RollNo: 3}, nil),
		}},
		{name: "Failure", args: args{rollNo: ""}, wantErr: errors.New("RollNo is not given"), mockCalls: []interface{}{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStudentService(mockdb, nil, nil)
			got, err := s.GetValidation(tt.args.rollNo)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StudentService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StudentService.GetValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMockstudentdatastore(ctrl)
	mockEnrollmentService := NewMockenrollmentServiceSample(ctrl)
	mockSubjectService := NewMocksubjectServiceSample(ctrl)
	tests := []struct {
		name      string
		s         models.Student
		mockCalls []interface{}
		wantErr   error
		want      string
	}{
		{name: "Valid Arguments", s: models.Student{Name: "Ankit", Age: 21, RollNo: 3}, mockCalls: []interface{}{
			mockdb.EXPECT().InsertStudent(gomock.Any()).Return(nil),
		}},
		{name: "Invalid Arguments", s: models.Student{Name: "", Age: 21, RollNo: 0}, wantErr: errors.New("RollNo and Name are mandatory")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStudentService(mockdb, mockEnrollmentService, mockSubjectService)
			err := s.PostValidation(tt.s)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StudentService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestStudentEnroll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMockstudentdatastore(ctrl)
	mockEnrollmentService := NewMockenrollmentServiceSample(ctrl)
	mockSubjectService := NewMocksubjectServiceSample(ctrl)
	tests := []struct {
		name      string
		s         models.Student
		id        int
		rollNo    int
		mockCalls []interface{}
		wantErr   error
		want      string
	}{
		{name: "Successful Creation of record", id: 1, rollNo: 1, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent(gomock.Any()).Return(models.Student{Name: "Ankit", Age: 21, RollNo: 1}, nil),
			mockSubjectService.EXPECT().GetValidation(gomock.All()).Return(models.Subject{Name: "Science", Id: 1}, nil),
			mockEnrollmentService.EXPECT().Insert(gomock.Any()).Return(nil),
		}},
		{name: "Student is not valid", s: models.Student{Name: "Ankit", Age: 21, RollNo: 1}, id: 1, rollNo: 1, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent(gomock.Any()).Return(models.Student{Name: "Ankit", Age: 21, RollNo: 1}, errors.New("Student doesn't found")),
		}, wantErr: errors.New("Student doesn't found")},

		{name: "Subject is not valid", s: models.Student{Name: "Ankit", Age: 21, RollNo: 1}, id: 1, rollNo: 1, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent(gomock.Any()).Return(models.Student{Name: "Ankit", Age: 21, RollNo: 1}, nil),
			mockSubjectService.EXPECT().GetValidation(gomock.Any()).Return(models.Subject{}, errors.New("Subject doesn't exist")),
		}, wantErr: errors.New("Subject doesn't exist")},

		{name: "Record Insertion Failed", s: models.Student{Name: "Ankit", Age: 21, RollNo: 1}, id: 1, rollNo: 1, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent(gomock.Any()).Return(models.Student{Name: "Ankit", Age: 21, RollNo: 1}, nil),
			mockSubjectService.EXPECT().GetValidation(gomock.All()).Return(models.Subject{Name: "Science", Id: 1}, nil),
			mockEnrollmentService.EXPECT().Insert(gomock.Any()).Return(errors.New("Record Insertion Failed")),
		}, wantErr: errors.New("Record Insertion Failed")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStudentService(mockdb, mockEnrollmentService, mockSubjectService)
			err := s.Enroll(tt.id, tt.rollNo)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StudentService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetSubs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMockstudentdatastore(ctrl)
	mockEnrollmentService := NewMockenrollmentServiceSample(ctrl)
	mockSubjectService := NewMocksubjectServiceSample(ctrl)
	tests := []struct {
		name      string
		rollNo    string
		mockCalls []interface{}
		wantErr   error
		want      []string
	}{
		{name: "Successful Got Records", rollNo: "1", mockCalls: []interface{}{
			mockEnrollmentService.EXPECT().GetSubs(gomock.Any()).Return([]int{1}, nil),
			mockSubjectService.EXPECT().GetValidation(1).Return(models.Subject{Name: "Science", Id: 1}, nil),
		}, want: []string{"Science"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStudentService(mockdb, mockEnrollmentService, mockSubjectService)
			studentNames, err := s.GetSubs(tt.rollNo)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StudentService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want, studentNames) {
				t.Errorf("StudentService.GetValidation() error = %v, wantErr %v", studentNames, tt.want)
				return
			}
		})
	}
}
