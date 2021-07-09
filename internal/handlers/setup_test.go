package handlers

import (
	"WebApp/internal/config"
	"WebApp/internal/models"
	"WebApp/internal/render"
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func getRoutes() http.Handler {
	//From main code
	// what I am going to put in the session
	gob.Register(models.ConfirmInfo{})

	// change this when in production
	app.InProduction = false

	session = scs.New()               // returns a new session manager with the default options. It is safe for concurrent use.
	session.Lifetime = 24 * time.Hour // session set for 24 hours
	session.Cookie.Persist = true     // cookie is persistent
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session // stores the session in our app-wide configuration.

	ts, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = ts
	app.UseCache = true

	repo := NewRepo(&app)
	NewHandlers(repo)
	render.NewTemplates(&app)

	// From routes code
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // middleware to handle panics gracefully
	//mux.Use(NoSurf)               // NoSurf used to combat CSRF attacks
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/coffee", Repo.Coffee)
	mux.Get("/cassava-cake", Repo.CassavaCake)

	mux.Get("/order", Repo.Order)
	mux.Get("/OrderAvailabilityJSON", Repo.OrderAvailabilityJSON)

	mux.Get("/confirm", Repo.Confirm)
	mux.Post("/confirm", Repo.PostConfirm)
	mux.Get("/order-summary", Repo.OrderSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}

// From middleware code
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next) // Constructs a new CSRFHandler that calls the specified handler if the CSRF check succeeds.

	// Sets the base cookie to use when building a CSRF token cookie.
	//This way you can specify the Domain, Path, HttpOnly, Secure, etc.
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session on every request.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// From the render code
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
