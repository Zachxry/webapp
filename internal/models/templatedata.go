package models

import "WebApp/internal/forms"

// TemplateData holds data sent from handlers to templates, makes availabe to every page.
type TemplateData struct {
	StringMap map[string]string      // map with a [key]value pair of [string]string
	IntMap    map[string]int         // map with a [key]value pair of [string]int
	FloatMap  map[string]float32     // map with a [key]value pair of [string]float32
	Data      map[string]interface{} // map with a [key]value pair of [string]interface{}
	// We use interface{} here when the data to be sent is unknown
	CSRFToken string // security token
	Flash     string // used for flash messages
	Warning   string // used for warnings
	Error     string // used for errors
	Form      *forms.Form
}
