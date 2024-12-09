package task10

import (
	"net/http"
)

func NewMuxer() *http.ServeMux {
	httpMux := http.NewServeMux()

	httpMux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	httpMux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	return httpMux
}
