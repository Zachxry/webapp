package handlers

import (
	"WebApp/pkg/config"
	"WebApp/pkg/models"
	"WebApp/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository structure contains a variable App which is a pointer to our config file
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr //RemoteAddr allows HTTP servers and other software to record the network address that sent the request, usually for logging.
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.Template(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip") // m.App.Session is used to access the SessionManager package functions.
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.Template(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Coffee renders the coffee page
func (m *Repository) Coffee(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "coffee.page.html", &models.TemplateData{})
}

// CassavaCake renders the cassava-cake page
func (m *Repository) CassavaCake(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "cassava-cake.page.html", &models.TemplateData{})
}

func (m *Repository) Order(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "order.page.html", &models.TemplateData{})
}

func (m *Repository) Confirm(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "confirm.page.html", &models.TemplateData{})
}
