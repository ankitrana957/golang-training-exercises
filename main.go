package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/student-api/handler"
	"github.com/student-api/stores/student"
)

func connection() (student.Sqldb, error) {
	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	if err != nil {
		return student.Sqldb{}, err
	}
	return student.Sqldb{con}, nil
}

func main() {
	con, err := connection()
	if err != nil {
		return
	}
	g := handler.CreateHandler(con)
	server := mux.NewRouter()
	server.HandleFunc("/student", g.GetAll).Methods("GET")
	server.HandleFunc("/student", g.Create).Methods("POST")
	server.HandleFunc("/student/{id}", g.Get).Methods("GET")
	server.HandleFunc("/student/{id}", g.Delete).Methods("DELETE")
	server.HandleFunc("/student/{id}", g.Update).Methods("PUT")
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8000", server))
}
