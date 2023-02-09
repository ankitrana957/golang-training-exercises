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
	conn1, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/student") // why?

	if err != nil {
		return
	}

	subjectService := service.NewSubStore(con)
	subjectHandler := handler.NewSubHandler(subjectService)

	enrollmentService := service.NewEnrollmentStore(con)
	studentService := service.NewStudentService(stores.SqlDb{conn1}, enrollmentService, subjectService)
	studentHandler := handler.NewStudentHandler(studentService)

	server := mux.NewRouter()

	server.HandleFunc("/subject", subjectHandler.GetSubject).Methods("GET")
	server.HandleFunc("/subject", subjectHandler.CreateSubject).Methods("POST")

	// server.HandleFunc("/student/", s.GetAll).Methods("GET")
	server.HandleFunc("/student", studentHandler.Create).Methods("POST")
	server.HandleFunc("/student", studentHandler.Get).Methods("GET")
	server.HandleFunc("/student/{rollNo}/subject/{id}", studentHandler.EnrollStudent).Methods("POST")

	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8000", server))
}
