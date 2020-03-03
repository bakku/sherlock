package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"bakku.org/sherlock/web/routes"
)

func TestGETHome_ShouldRenderSuccessfully(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("creating a new request should be nil")
	}

	rr := httptest.NewRecorder()
	handler := routes.Home{}
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatal("expected error code to be 200")
	}

	if rr.Body.String() != "Hello, World" {
		t.Fatal("expected body to contain 'Hello, World'")
	}
}
