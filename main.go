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

func connection() (stores.SqlDb, error) {
	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	if err != nil {
		return stores.SqlDb{}, err
	}
	return stores.SqlDb{con}, nil
}

func main() {
	con, err := connection()
	con1, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	if err != nil {
		return
	}

	subjectService := service.NewSubStore(con)
	subjectHandler := handler.NewSubHandler(subjectService)

	enrollmentService := service.NewEnrollmentStore(con)
	studentService := service.NewStudentService(stores.SqlDb{con1}, enrollmentService, subjectService)
	studentHandler := handler.NewStudentHandler(studentService)

	server := mux.NewRouter()

	server.HandleFunc("/subject", subjectHandler.GetSubject).Methods("GET")
	server.HandleFunc("/subject", subjectHandler.InsertSubject).Methods("POST")
	server.HandleFunc("/student", studentHandler.InsertStudent).Methods("POST")
	server.HandleFunc("/student", studentHandler.GetStudent).Methods("GET")
	server.HandleFunc("/student/{rollNo}/subject/{id}", studentHandler.EnrollStudent).Methods("POST")

	// List all subjects of a particular student
	server.HandleFunc("/student/{rollNo}/subject", studentHandler.GetAllSubs).Methods("GET")

	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8000", server))
}
