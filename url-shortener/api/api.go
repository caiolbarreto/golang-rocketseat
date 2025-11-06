package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/shorten", handlePost(db))
	r.Get("/{shortCode}", handleGet(db))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type ResponseBody struct {
	Error string `json:"error,omitempty"`
	Data  string `json:"data,omitempty"`
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ResponseBody{
				Error: "invalid request body",
			})
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ResponseBody{
				Error: "invalid URL",
			})
			return
		}

		code := genCode()
		db[code] = body.URL
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ResponseBody{
			Data: code,
		})
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func genCode() string {
	byts := make([]byte, 8)
	for i := range byts {
		byts[i] = charset[rand.Intn(len(charset))]
	}
	return string(byts)
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortCode := chi.URLParam(r, "shortCode")
		url, ok := db[shortCode]

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ResponseBody{
				Error: "short code not found",
			})
			return
		}

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
