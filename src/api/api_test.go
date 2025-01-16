package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func Test_Ping(t *testing.T) {
	api := New(&Opts{})
	t.Run("Test Ping", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/ping", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.GET("/ping", api.Ping)

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := "pong"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("Test Ping with different method", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/ping", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.GET("/ping", api.Ping)

		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})
}

func Test_Self(t *testing.T) {
	api := New(&Opts{})
	t.Run("Test Self", func(t *testing.T) {
		body := strings.NewReader(`{"key":"value"}`)
		req, err := http.NewRequest("POST", "/self", body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.POST("/self", api.Self)

		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"key":"value"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("Test Self with different method", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/self", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := httprouter.New()
		router.POST("/self", api.Self)

		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})
}
