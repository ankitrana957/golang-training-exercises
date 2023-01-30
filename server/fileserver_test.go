package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestGetPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	c := handler{db}
	tests := []struct {
		name string
		r    *http.Request
		w    *httptest.ResponseRecorder
		exp  string
	}{
		{name: "Get Request From Person", w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/student", nil), exp: `[{"Name":"Ankit Rana","Age":20,"Phn":"738213671"}]`},
		{name: "Request for ping", w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/ping", nil), exp: "Pong"},
		{name: "Wrong Url", w: httptest.NewRecorder(), r: httptest.NewRequest(http.MethodGet, "/sdsha", nil), exp: "Wrong Url"},
	}

	rs := mock.NewRows([]string{"name", "age", "phn"}).AddRow("Ankit Rana", 20, "738213671")
	mock.ExpectQuery("SELECT *").WillReturnRows(rs)

	for _, tt := range tests {
		c.GetData(tt.w, tt.r)
		res := tt.w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Expected error to be nil got %v", err)
		}
		output := string(data)
		n := strings.TrimSpace(output)

		if n != tt.exp {
			t.Errorf("Expected %v got %v", strings.TrimSpace(tt.exp), strings.TrimSpace(output))
		}
	}
}
