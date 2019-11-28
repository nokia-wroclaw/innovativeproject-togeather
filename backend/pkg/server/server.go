package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	// smth like config goes here
	router chi.Router
}

func New() *Server {
	s := &Server{}

	r := chi.NewRouter()

	r.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Timeout(10*time.Second),
	)

	// add handlers
	pingHandler := pingHandler{}

	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", pingHandler.ping)
	})

	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(jsonData)
}

func respondError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
