package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//AppConfig holds the application config
type AppConfig struct {
	UseCache      bool                          // sets value
	TemplateCache map[string]*template.Template // creates a map[string] pointer to html/template package.Template structure
	InfoLog       *log.Logger                   // logs relevant information as needed
	ErrorLog      *log.Logger                   // logs relevant error information as needed
	InProduction  bool                          // lets us set values T for when in production
	Session       *scs.SessionManager           // allows for control of session length
}
