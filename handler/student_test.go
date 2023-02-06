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

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

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

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	vars := map[string]string{
		"id": "8",
	}
	mockdatastore := NewMockdatastore(ctrl)
	h := studentHandler{db: mockdatastore}
	s := models.Student{Name: "Ankit", Age: 21, RollNo: 2}
	test := []struct {
		name     string
		w        *httptest.ResponseRecorder
		r        *http.Request
		mockCall []interface{}
		want     string
	}{
		{name: "Student Found", w: httptest.NewRecorder(), r: mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/student", nil), vars), want: `{"Name":"Ankit","Age":21,"RollNo":2}`, mockCall: []interface{}{
			mockdatastore.EXPECT().Get(vars["id"]).Return(s, nil),
		}},
		{name: "Student Doesn't Found", w: httptest.NewRecorder(), r: mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/student", nil), vars), want: `Got an error`, mockCall: []interface{}{
			mockdatastore.EXPECT().Get(vars["id"]).Return(models.Student{}, errors.New("Got an error")),
		}},
	}
	for _, tt := range test {
		h.Get(tt.w, tt.r)
		result := getRequestResponse(*tt.w)
		if !reflect.DeepEqual(tt.want, result) {
			t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
		}
	}
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	h := studentHandler{
		db: mockdatastore,
	}
	student := []models.Student{
		{Name: "Ankit", Age: 21, RollNo: 1},
		{Name: "Amit", Age: 21, RollNo: 2},
	}
	test := []struct {
		name     string
		w        *httptest.ResponseRecorder
		r        *http.Request
		mockCall []interface{}
		want     string
	}{
		{name: "Students Found", w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/student", nil), want: `[{"Name":"Ankit","Age":21,"RollNo":1},{"Name":"Amit","Age":21,"RollNo":2}]`, mockCall: []interface{}{
			mockdatastore.EXPECT().GetAll().Return(student, nil),
		}},
		{name: "Students Doesn't Found", w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/student", nil), want: `Error in getting response`, mockCall: []interface{}{
			mockdatastore.EXPECT().GetAll().Return(student, errors.New("Error in getting response")),
		}},
	}

	for _, tt := range test {
		h.GetAll(tt.w, tt.r)
		result := getRequestResponse(*tt.w)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
		}
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)

	h := studentHandler{db: mockdatastore}
	test := []struct {
		name     string
		w        *httptest.ResponseRecorder
		r        *http.Request
		mockCall []interface{}
		want     string
	}{
		{
			name: "Student Inserted", w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(`{"Name":"Ankit","Age":21,"RollNo":1}`))), want: `Succesfully Inserted`, mockCall: []interface{}{
				mockdatastore.EXPECT().Insert(models.Student{Name: "Ankit", Age: 21, RollNo: 1}).Return(nil),
			}},
		{
			name: "Invalid JSON Format", w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(`{"Name":"Ankit","Age":21,"RollNo":1`))), want: `Invalid JSON Format`, mockCall: []interface{}{},
		},
		{
			name: "Error Occured", w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(`{"Name":"Ankit","Age":21,"RollNo":1}`))), want: `Some Error`, mockCall: []interface{}{
				mockdatastore.EXPECT().Insert(models.Student{Name: "Ankit", Age: 21, RollNo: 1}).Return(errors.New("Some Error")),
			},
		},
	}
	for _, tt := range test {
		h.Create(tt.w, tt.r)
		result := getRequestResponse(*tt.w)
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
		}
	}

}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodDelete, "/student", nil), vars)
	mockdatastore.EXPECT().Delete(vars["id"]).Return(nil)
	h := studentHandler{db: mockdatastore}
	h.Delete(w, r)

	want := "Successfully Deleted"
	result := getRequestResponse(*w)

	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestDeleteErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodDelete, "/student", nil), vars)
	mockdatastore.EXPECT().Delete(vars["id"]).Return(errors.New("Deletion Failed"))
	h := studentHandler{db: mockdatastore}
	h.Delete(w, r)

	want := "Deletion Failed"
	result := getRequestResponse(*w)

	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	payload := `{"Name":"Ankit","Age":21}`
	want := `SuccessFull Put Request`
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/student", bytes.NewBuffer([]byte(payload))), vars)
	mockdatastore.EXPECT().Update(models.Student{Name: "Ankit", Age: 21, RollNo: 8}).Return(nil)
	h := studentHandler{
		db: mockdatastore,
	}
	h.Update(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestUpdateErr1(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	payload := `{"Name":"Ankit","Age":21`
	want := `Invalid JSON Format`
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/student", bytes.NewBuffer([]byte(payload))), vars)
	h := studentHandler{
		db: mockdatastore,
	}
	h.Update(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestUpdateErr2(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	payload := `{"Name":"Ankit","Age":21}`
	want := `Updation Failed`
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/student", bytes.NewBuffer([]byte(payload))), vars)
	mockdatastore.EXPECT().Update(models.Student{Name: "Ankit", Age: 21, RollNo: 8}).Return(errors.New("Updation Failed"))
	h := studentHandler{
		db: mockdatastore,
	}
	h.Update(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}
