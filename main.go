package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/student-api/handler"
	"github.com/student-api/service"
	"github.com/student-api/stores"
)

func studentDb() (stores.StudentStore, error) {
	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	if err != nil {
		return stores.StudentStore{}, err
	}
	return stores.StudentStore{con}, nil
}

func subjectDb() (stores.SubjectStore, error) {
	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	if err != nil {
		return stores.SubjectStore{}, err
	}
	return stores.SubjectStore{con}, nil
}

func enrollmentDb() (stores.EnrollmentStore, error) {
	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	if err != nil {
		return stores.EnrollmentStore{}, err
	}
	return stores.EnrollmentStore{con}, nil
}

func main() {
	con1, err1 := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	con2, err2 := subjectDb()
	con3, err3 := enrollmentDb()
	if err1 != nil && err2 != nil && err3 != nil {
		return
	}

	subjectService := service.NewSubStore(con2)
	subjectHandler := handler.NewSubHandler(subjectService)

	enrollmentService := service.NewEnrollmentStore(con3)
	studentService := service.NewStudentService(stores.StudentStore{con1}, enrollmentService, subjectService)
	studentHandler := handler.NewStudentHandler(studentService)

	server := mux.NewRouter()

	server.HandleFunc("/subject", subjectHandler.GetSubject).Methods("GET")
	server.HandleFunc("/subject", subjectHandler.InsertSubject).Methods("POST")
	server.HandleFunc("/student", studentHandler.InsertStudent).Methods("POST")
	server.HandleFunc("/student", studentHandler.GetStudent).Methods("GET")

	// Enroll student with subject
	server.HandleFunc("/student/{rollNo}/subject/{id}", studentHandler.EnrollStudent).Methods("POST")

	// List all subjects of a particular student
	server.HandleFunc("/student/{rollNo}/subject", studentHandler.GetAllSubs).Methods("GET")

	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8000", server))
}
