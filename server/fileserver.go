package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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

type config struct {
	user     string
	password string
	host     string
	port     string
	database string
}

var cred = config{
	user:     "root",
	password: "root",
	host:     "localhost",
	port:     "3306",
	database: "student",
}

func establishConnection(driverName, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	return &(*db)
}

func getDataSourceName(cred config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cred.user, cred.password, cred.host, cred.port, cred.database)
}

func retrievingData(c *sql.DB) (s []person, err error) {
	rows, err := c.Query("SELECT * FROM personDetails")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var stud person
		err := rows.Scan(&stud.Name, &stud.Age, &stud.Phn)
		if err != nil {
			return nil, err
		}
		s = append(s, stud)
	}
	return s, nil
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/get/student" {
		src := getDataSourceName(cred)
		db := establishConnection("mysql", src)
		data, _ := retrievingData(db)
		json.NewEncoder(w).Encode(data)
	} else if r.URL.Path == "/ping" {
		w.Write([]byte("Pong"))
	} else {
		b := fmt.Sprintf("Wrong Url %s", r.URL.Path)
		w.Write([]byte(b))
	}
}

func StartServer() {
	log.Fatal(http.ListenAndServe(":8000", http.HandlerFunc(getStudent)))
}
