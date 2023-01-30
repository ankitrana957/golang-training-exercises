package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Function Closure
// func getcount() func(http.ResponseWriter, *http.Request) {
// 	a := 0
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.URL.Path == "/hello" {
// 			s := fmt.Sprintf("Hey you type hello %d", a)
// 			w.Write([]byte(s))
// 			a += 1
// 		} else {
// 			b := fmt.Sprintf("Hello %s", r.URL.Path)
// 			w.Write([]byte(b))
// 		}
// 	}
// }

type person struct {
	Name string
	Age  int
	Phn  string
}

type handler struct {
	*sql.DB
}

func connectToDatabase() handler {
	database, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/student")
	return handler{database}
}

func (h handler) GetData(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/student":
		s := []person{}
		rows, err := h.Query("SELECT * FROM personDetails")
		if err != nil {
			return
		}
		defer rows.Close()
		for rows.Next() {
			var stud person
			rows.Scan(&stud.Name, &stud.Age, &stud.Phn)
			s = append(s, stud)
		}
		json.NewEncoder(w).Encode(s)
	case "/ping":
		s := "Pong"
		w.Write([]byte(s))
		return
	default:
		s := "Wrong Url"
		w.Write([]byte(s))
		return
	}
	w.WriteHeader(200)
}
