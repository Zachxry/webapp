package main

import (
	"WebApp/internal/config"
	"WebApp/internal/handlers"
	"WebApp/internal/models"
	"WebApp/internal/render"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig        // // creates a variable app which is has the value of the config structure
var session *scs.SessionManager // creates a variable session which is a pointer to SessionManager structure

// main is the main application function
func main() {

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

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNumber)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
