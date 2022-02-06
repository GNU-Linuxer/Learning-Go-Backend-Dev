package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "txt")
		w.Write([]byte("welcome"))
	})

	// RESTy routes for "articles" resource
	r.Route("/articles", func(r chi.Router) {
		// GET /articles
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "txt")
			w.Write([]byte("List Article Page"))
		})
		// GET /articles/01-16-2017
		r.Get("/{month}-{day}-{year}", func(w http.ResponseWriter, r *http.Request) {
			responseText := "Listing Articles on " + chi.URLParam(r, "month") + "-" + chi.URLParam(r, "day") + "-" + chi.URLParam(r, "year")
			w.Write([]byte(responseText))
		})
		// POST /articles
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			received := r.Body

			w.Write([]byte(received))
		})
	})

	http.ListenAndServe(":9000", r)
}
