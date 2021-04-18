package main

import (
	"WebApp/internal/config"
	"WebApp/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes handles the routing of the pages
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // middleware to handle panics gracefully
	mux.Use(NoSurf)               // NoSurf used to combat CSRF attacks
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/coffee", handlers.Repo.Coffee)
	mux.Get("/cassava-cake", handlers.Repo.CassavaCake)

	mux.Get("/order", handlers.Repo.Order)
	mux.Post("/order", handlers.Repo.PostOrder)
	mux.Get("/OrderAvailabilityJSON", handlers.Repo.OrderAvailabilityJSON)

	mux.Get("/confirm", handlers.Repo.Confirm)
	mux.Post("/confirm", handlers.Repo.PostConfirm)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
