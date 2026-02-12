package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	// "github.com/smithwithatypo/WeLearnAnything/internal/handlers"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", healthHandler)
		// ... other API routes go here
	})
	// Routes
	// r.Post("/api/summarize", handlers.SummarizeJobDescription)   # TODO:  update routes

	r.Handle("/*", http.FileServer(http.Dir("../frontend/dist")))

	http.ListenAndServe(":8080", r)
}
