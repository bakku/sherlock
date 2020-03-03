package routes

import "net/http"

type Home struct{}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
