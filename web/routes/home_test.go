package routes_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bakku.org/sherlock/web/routes"
)

func TestGETHome_ShouldRenderSuccessfully(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal("creating a new request should be nil")
	}

	rr := httptest.NewRecorder()
	handler := routes.Home{Template: getTemplate("home.html")}
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatal("expected error code to be 200")
	}

	if !strings.Contains(rr.Body.String(), "Home") {
		t.Fatal("expected body to contain 'Home'")
	}
}
