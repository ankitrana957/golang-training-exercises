package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/student-api/handler"
	"github.com/student-api/stores"
)

func connection() (stores.Sqldb, error) {
	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	if err != nil {
		return stores.Sqldb{}, err
	}
	return stores.Sqldb{con}, nil
}

func main() {
	con, err := connection()
	if err != nil {
		return
	}
	g := handler.CreateAHandler(con)
	server := mux.NewRouter()
	server.HandleFunc("/student", g.GetAll).Methods("GET")
	server.HandleFunc("/student", g.Create).Methods("POST")
	server.HandleFunc("/student/{id}", g.Get).Methods("GET")
	server.HandleFunc("/student/{id}", g.Delete).Methods("DELETE")
	server.HandleFunc("/student/{id}", g.Update).Methods("PUT")
	http.Handle("/", server)
	http.ListenAndServe(":8000", server)
}
