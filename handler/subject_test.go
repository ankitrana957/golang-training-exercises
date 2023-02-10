package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func TestCreateSubject(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSubHandler := NewMocksubjectHandler(ctrl)
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	sub := models.Subject{Name: "Science", Id: 1}
	tests := []struct {
		name      string
		serv      subjectHandler
		args      args
		mockCalls []interface{}
		want      string
	}{
		{
			name: "Insert Subject", serv: mockSubHandler, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/subject", bytes.NewBuffer([]byte(`{"Name":"Science","Id":1}`)))}, mockCalls: []interface{}{
				mockSubHandler.EXPECT().InsertValidation(sub).Return(nil),
			}, want: "Successfully inserted subject",
		},
		{
			name: "Insertion Fail", serv: mockSubHandler, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/subject", bytes.NewBuffer([]byte(`{"Name":"Science","Id":1}`)))}, mockCalls: []interface{}{
				mockSubHandler.EXPECT().InsertValidation(sub).Return(errors.New("Failed to insert subject")),
			}, want: "Failed to insert subject",
		},
		{
			name: "Invalid Json", serv: mockSubHandler, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/subject", bytes.NewBuffer([]byte(`{"Name":"Science","Id":1`)))}, want: "Invalid JSON format",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := subService{
				serv: tt.serv,
			}
			s.InsertSubject(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}

func TestGetSubject(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockSubHandler := NewMocksubjectHandler(ctrl)
	sub := models.Subject{Name: "Science", Id: 1}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name      string
		serv      subjectHandler
		args      args
		mockCalls []interface{}
		want      string
	}{
		{name: "Successful Retrieval", serv: mockSubHandler, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/subject?id=1", nil)}, mockCalls: []interface{}{
			mockSubHandler.EXPECT().GetValidation(1).Return(sub, nil),
		}, want: `{"Id":1,"Name":"Science"}`,
		},
		{name: "Subject Retrieval Failed", serv: mockSubHandler, args: args{w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/subject?id=1", nil)}, mockCalls: []interface{}{
			mockSubHandler.EXPECT().GetValidation(1).Return(models.Subject{}, errors.New("Subject doesn't exist")),
		}, want: `Subject doesn't exist`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := subService{
				serv: tt.serv,
			}
			s.GetSubject(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}
