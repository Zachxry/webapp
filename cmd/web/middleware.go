package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf provides a CSRFHandler that wraps the http.Handler and checks for CSRF attacks
// on every non-safe (non-GET/HEAD/OPTIONS/TRACE) method.
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
