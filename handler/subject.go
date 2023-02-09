package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/student-api/models"
)

type subjectHandler interface {
	GetValidation(id int) (models.Subject, error)
	InsertValidation(sub models.Subject) error
}

type subService struct {
	serv subjectHandler
}

func NewSubHandler(serv subjectHandler) subService {
	return subService{serv}
}

func (s subService) CreateSubject(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)

	sub := models.Subject{}
	err1 := json.Unmarshal(data, &sub)
	if err1 != nil {

		fmt.Fprintln(w, "Invalid JSON format")
		return
	}
	err2 := s.serv.InsertValidation(sub)
	if err2 != nil {
		fmt.Fprintln(w, err2.Error())
		return
	}
	fmt.Fprintln(w, "Successfully inserted subject")
}

func (s subService) GetSubject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	newId, _ := strconv.Atoi(id)
	v, err := s.serv.GetValidation(newId)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	json.NewEncoder(w).Encode(v)
}
