package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/student-api/models"
)

type studentEnrollmentService interface {
	GetValidation(rollNo string) (models.Student, error)
	PostValidation(models.Student) error
	Enroll(id, rollNo int) error
	GetSubs(rollNo string) ([]string, error)
}

type serviceHandler struct {
	serv studentEnrollmentService
}

// Factory
func NewStudentHandler(serv studentEnrollmentService) serviceHandler {
	return serviceHandler{serv}
}

func (h serviceHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	rollNo := r.URL.Query().Get("rollNo")
	s, err := h.serv.GetValidation(rollNo)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	json.NewEncoder(w).Encode(s)

}

func (h serviceHandler) InsertStudent(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	s := models.Student{}
	err1 := json.Unmarshal(data, &s)
	if err1 != nil {
		fmt.Fprintln(w, "Invalid JSON Format")
		return
	}
	err2 := h.serv.PostValidation(s)
	if err2 != nil {
		fmt.Fprint(w, err2)
		return
	}
	fmt.Fprint(w, "Successfully Inserted")
}

func (s serviceHandler) EnrollStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	rollNo, _ := strconv.Atoi(params["rollNo"])
	err := s.serv.Enroll(id, rollNo)
	if err != nil {
		fmt.Fprint(w, "Failed to Enroll Student")
		return
	}
	fmt.Fprint(w, "Succesfully Enrolled Student with Subject")
}

func (s serviceHandler) GetAllSubs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rollNo, _ := params["rollNo"]
	res, err := s.serv.GetSubs(rollNo)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, res)
}
