package handler

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"

	models "github.com/student-api/models"
)

func getRequestResponse(w httptest.ResponseRecorder) (result string) {
	res := w.Result()
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	formattedData := string(data)
	result = strings.TrimSpace(formattedData)
	return
}

func TestStudentGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentEnrollmentService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	s := models.Student{Name: "Ankit", Age: 21, RollNo: 2}

	tests := []struct {
		name     string
		serv     studentEnrollmentService
		args     args
		mockCall []interface{}
		want     string
	}{
		{name: "Student Found", serv: mockStudentService, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/student?rollNo=1", nil)}, mockCall: []interface{}{
			mockStudentService.EXPECT().GetValidation(gomock.All()).Return(s, nil),
		}, want: `{"Name":"Ankit","Age":21,"RollNo":2}`,
		},
		{name: "Student Not Found", serv: mockStudentService, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/student?rollNo=1", nil)}, mockCall: []interface{}{
			mockStudentService.EXPECT().GetValidation(gomock.All()).Return(models.Student{}, errors.New("Student doesn't exist")),
		}, want: `Student doesn't exist`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := serviceHandler{
				serv: tt.serv,
			}
			h.GetStudent(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}

func TestStudentCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentEnrollmentService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}

	tests := []struct {
		name      string
		serv      studentEnrollmentService
		args      args
		mockCalls []interface{}
		want      string
	}{
		{
			name: "Successfully Inserted", serv: mockStudentService, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(`{"Name":"Ankit","Age":21,"RollNo":1}`)))},
			mockCalls: []interface{}{
				mockStudentService.EXPECT().PostValidation(models.Student{Name: "Ankit", Age: 21, RollNo: 1}).Return(nil),
			},
			want: "Successfully Inserted",
		},
		{
			name: "Invalid Json Format", serv: mockStudentService, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(`{"Name":"Ankit","Age":21,"RollNo":1`)))},
			want: "Invalid JSON Format",
		},
		{
			name: "Insertion Failed", serv: mockStudentService, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(`{"Name":"Ankit","Age":21,"RollNo":1}`)))},
			mockCalls: []interface{}{
				mockStudentService.EXPECT().PostValidation(models.Student{Name: "Ankit", Age: 21, RollNo: 1}).Return(errors.New("Failed to insert data")),
			},
			want: "Failed to insert data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := serviceHandler{
				serv: tt.serv,
			}
			h.InsertStudent(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}

func TestEnrollStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentEnrollmentService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		serv     studentEnrollmentService
		args     args
		mockCall []interface{}
		want     string
	}{
		{
			name: "Succesfully Enrolled Student", serv: mockStudentService, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/student/1/subject/2", nil)}, mockCall: []interface{}{
				mockStudentService.EXPECT().Enroll(gomock.Any(), gomock.Any()).Return(nil),
			},
			want: "Succesfully Enrolled Student with Subject",
		},
		{
			name: "Fail to Enroll Student", serv: mockStudentService, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/student/1/subject/2", nil)}, mockCall: []interface{}{
				mockStudentService.EXPECT().Enroll(gomock.Any(), gomock.Any()).Return(errors.New("Failed to Enroll Student")),
			},
			want: "Failed to Enroll Student",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := serviceHandler{
				serv: tt.serv,
			}
			s.EnrollStudent(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}

func TestGetAllSubs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentEnrollmentService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name      string
		serv      studentEnrollmentService
		args      args
		mockCalls []interface{}
		want      string
	}{
		{name: "Get Subjects", serv: mockStudentService, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/student/1/subject", nil)}, mockCalls: []interface {
		}{
			mockStudentService.EXPECT().GetSubs(gomock.Any()).Return([]string{"English", "Science"}, nil),
		}, want: `[English Science]`},
		{name: "Failed to get Student", serv: mockStudentService, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/student/1/subject", nil)}, mockCalls: []interface {
		}{
			mockStudentService.EXPECT().GetSubs(gomock.Any()).Return(nil, errors.New("Student Doesn't exist")),
		}, want: `Student Doesn't exist`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := serviceHandler{
				serv: tt.serv,
			}
			s.GetAllSubs(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}
