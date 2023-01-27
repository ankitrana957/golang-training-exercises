package api

import (
	"fmt"
	"net/http"
	"strings"
)

func UserRecord() func(http.ResponseWriter, *http.Request) {
	var v = map[string]int{}
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if url[:4] == "/get" {
			b := strings.Split(url, "/")
			l := len(b)
			name := b[l-1]
			_, exist := v[name]
			if exist == true {
				e := fmt.Sprintf("Welcome back %s !", name)
				w.Write([]byte(e))
			} else {
				e := fmt.Sprintf("Hello %s", name)
				w.Write([]byte(e))
				v[name] = 1
			}
		} else {
			w.Write([]byte("Wrong Path"))
		}
	}
}

func StartServer() {
	http.ListenAndServe(":8000", http.HandlerFunc(UserRecord()))
}
