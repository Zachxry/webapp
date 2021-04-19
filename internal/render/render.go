package render

import (
	"WebApp/internal/config"
	"WebApp/internal/models"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td // returns template data struct, we can add data here as needed
}

// Template is used to render our html templates
func Template(w http.ResponseWriter, r *http.Request, html string, td *models.TemplateData) {

	var tc map[string]*template.Template // variable tc that is a map with a [key]value pair of [string] pointer to template.Template
	// Template is a specialized Template from "text/template" that produces a safe HTML document fragment.
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[html]
	if !ok {
		log.Fatal("could not get template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r) // adds data we would want to have on pages by default as needed

	_ = t.Execute(buf, td) // where our data gets passed from the pointer to models.TemplateData

	_, err := buf.WriteTo(w) // write the bytes of data to the browser
	if err != nil {
		fmt.Println("Error writing to browser", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
