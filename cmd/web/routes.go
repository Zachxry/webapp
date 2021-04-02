package main

import (
	"WebApp/pkg/config"
	"WebApp/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

// routes handles the routing of the pages
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // middleware to handle panics gracefully
	mux.Use(NoSurf) // NoSurf used to combat CSRF attacks
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
	}


