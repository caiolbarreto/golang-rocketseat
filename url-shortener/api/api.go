package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/shorten", handlePost)
	r.Get("/{shortCode}", handleGet)

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type ResponseBody struct {
	Error string `json:"error,omitempty"`
	Data  string `json:"data,omitempty"`
}


func handlePost(w http.ResponseWriter, r *http.Request) {
}

func handleGet(w http.ResponseWriter, r *http.Request) {
}