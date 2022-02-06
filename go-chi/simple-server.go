package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Starting Server...")
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) { // A handler always has a response writer and request
		w.Header().Set("Content-Type", "txt")
		w.Write([]byte("welcome"))
	})
	router.Get("/json", func(w http.ResponseWriter, r *http.Request) { // A handler always has a response writer and request
		w.Header().Set("Content-Type", "json")
		json.NewEncoder(w).Encode("JSON here!")
	})

	// RESTy routes for "articles" resource
	router.Route("/articles", func(r chi.Router) {
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
			//received := r.Body

			w.Write([]byte("Received post action on article"))
		})
	})

	http.ListenAndServe(":9000", router)
}
