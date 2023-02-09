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
	type fields struct {
		db         studentdatastore
		enrollment enrollmentServiceSample
		subject    subjectServiceSample
	}
	type args struct {
		rollNo string
	}
	s := models.Student{Name: "Ankit", Age: 21, RollNo: 3}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      models.Student
		mockCalls []interface{}
		wantErr   error
	}{
		{name: "Success", fields: fields{db: mockdb}, args: args{rollNo: "1"}, want: s, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent("1").Return(s, nil),
		}},
		{name: "Failure", fields: fields{db: mockdb}, args: args{rollNo: ""}, wantErr: errors.New("RollNo is not given"), mockCalls: []interface{}{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudentEnrollmentService{
				db:         tt.fields.db,
				enrollment: tt.fields.enrollment,
				subject:    tt.fields.subject,
			}
			got, err := s.GetValidation(tt.args.rollNo)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StudentEnrollmentService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StudentEnrollmentService.GetValidation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentEnrollmentService_PostValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMockstudentdatastore(ctrl)
	mockEnrollmentService := NewMockenrollmentServiceSample(ctrl)
	mockSubjectService := NewMocksubjectServiceSample(ctrl)
	type fields struct {
		db         studentdatastore
		enrollment enrollmentServiceSample
		subject    subjectServiceSample
	}
	tests := []struct {
		name      string
		fields    fields
		s         models.Student
		mockCalls []interface{}
		wantErr   error
		want      string
	}{
		{name: "Valid Arguments", fields: fields{
			db: mockdb, enrollment: mockEnrollmentService, subject: mockSubjectService,
		}, s: models.Student{Name: "Ankit", Age: 21, RollNo: 3}, mockCalls: []interface{}{
			mockdb.EXPECT().InsertStudent(gomock.Any()).Return(nil),
		}},
		{name: "Invalid Arguments", fields: fields{
			db: mockdb, enrollment: mockEnrollmentService, subject: mockSubjectService,
		}, s: models.Student{Name: "", Age: 21, RollNo: 0}, wantErr: errors.New("RollNo and Name are mandatory")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudentEnrollmentService{
				db:         tt.fields.db,
				enrollment: tt.fields.enrollment,
				subject:    tt.fields.subject,
			}
			err := s.PostValidation(tt.s)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StudentEnrollmentService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestStudentEnrollmentService_Enroll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := NewMockstudentdatastore(ctrl)
	mockEnrollmentService := NewMockenrollmentServiceSample(ctrl)
	mockSubjectService := NewMocksubjectServiceSample(ctrl)
	type fields struct {
		db         studentdatastore
		enrollment enrollmentServiceSample
		subject    subjectServiceSample
	}
	stud := models.Student{Name: "Ankit", Age: 21, RollNo: 1}
	sub := models.Subject{Name: "Science", Id: 1}
	tests := []struct {
		name      string
		fields    fields
		s         models.Student
		id        int
		rollNo    int
		mockCalls []interface{}
		wantErr   error
		want      string
	}{
		{name: "Successful Creation of record", fields: fields{db: mockdb, enrollment: mockEnrollmentService, subject: mockSubjectService}, s: models.Student{Name: "Ankit", Age: 21, RollNo: 1}, id: 1, rollNo: 1, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent(gomock.Any()).Return(stud, nil),
			mockSubjectService.EXPECT().GetValidation(gomock.All()).Return(sub, nil),
			mockEnrollmentService.EXPECT().Insert(gomock.Any()).Return(nil),
		}},
		{name: "Student is not valid", fields: fields{db: mockdb, enrollment: mockEnrollmentService, subject: mockSubjectService}, s: models.Student{Name: "Ankit", Age: 21, RollNo: 1}, id: 1, rollNo: 1, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent(gomock.Any()).Return(stud, errors.New("Student doesn't found")),
		}, wantErr: errors.New("Student doesn't found")},

		{name: "Subject is not valid", fields: fields{db: mockdb, enrollment: mockEnrollmentService, subject: mockSubjectService}, s: models.Student{Name: "Ankit", Age: 21, RollNo: 1}, id: 1, rollNo: 1, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent(gomock.Any()).Return(stud, nil),
			mockSubjectService.EXPECT().GetValidation(gomock.Any()).Return(models.Subject{}, errors.New("Subject doesn't exist")),
		}, wantErr: errors.New("Subject doesn't exist")},

		{name: "Record Insertion Failed", fields: fields{db: mockdb, enrollment: mockEnrollmentService, subject: mockSubjectService}, s: models.Student{Name: "Ankit", Age: 21, RollNo: 1}, id: 1, rollNo: 1, mockCalls: []interface{}{
			mockdb.EXPECT().GetStudent(gomock.Any()).Return(stud, nil),
			mockSubjectService.EXPECT().GetValidation(gomock.All()).Return(sub, nil),
			mockEnrollmentService.EXPECT().Insert(gomock.Any()).Return(errors.New("Record Insertion Failed")),
		}, wantErr: errors.New("Record Insertion Failed")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StudentEnrollmentService{
				db:         tt.fields.db,
				enrollment: tt.fields.enrollment,
				subject:    tt.fields.subject,
			}
			err := s.Enroll(tt.id, tt.rollNo)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StudentEnrollmentService.GetValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
