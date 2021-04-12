package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string      // map with a [key]value pair of [string]string
	IntMap    map[string]int         // map with a [key]value pair of [string]int
	FloatMap  map[string]float32     // map with a [key]value pair of [string]float32
	Data      map[string]interface{} // map with a [key]value pair of [string]interface{}
	// We use interface{} here when we don't know what data will be sent
	CSRFToken string // security token
	Flash     string // used for flash messages
	Warning   string // used for warnings
	Error     string // used for errors
}
