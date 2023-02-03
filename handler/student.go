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

type datastore interface {
	Insert(models.Student) error
	GetAll() ([]models.Student, error)
	Get(string) (models.Student, error)
	Delete(string) error
	Update(models.Student) error
}

func CreateAHandler(db datastore) studentHandler {
	return studentHandler{db}
}

type studentHandler struct {
	db datastore
}

func (h studentHandler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	s, err := h.db.Get(id)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	json.NewEncoder(w).Encode(s)
	return
}

func (h studentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	s, err := h.db.GetAll()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	json.NewEncoder(w).Encode(s)
}

func (h studentHandler) Create(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	s := models.Student{}
	err1 := json.Unmarshal(data, &s)
	if err1 != nil {
		fmt.Fprintln(w, "Invalid JSON Format")
		return
	}
	err2 := h.db.Insert(s)
	if err2 != nil {
		fmt.Fprint(w, err2)
		return
	}
	fmt.Fprint(w, "Succesfully Inserted")
}

func (h studentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	err := h.db.Delete(id)
	if err != nil {
		fmt.Fprint(w, "Deletion Failed")
		return
	}
	fmt.Fprint(w, "Successfully Deleted")
	return
}

func (h studentHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	data, _ := io.ReadAll(r.Body)
	newData := struct {
		Name string
		Age  int
	}{}
	err2 := json.Unmarshal(data, &newData)
	if err2 != nil {
		fmt.Fprintln(w, "Invalid JSON Format")
		return
	}
	convertedId, _ := strconv.Atoi(id)
	s := models.Student{
		Name:   newData.Name,
		Age:    newData.Age,
		RollNo: convertedId,
	}
	err3 := h.db.Update(s)
	if err3 != nil {
		fmt.Fprintln(w, "Updation Failed")
		return
	}
	fmt.Fprintln(w, "SuccessFull Put Request")
	return
}
