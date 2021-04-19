package handlers

import (
	"WebApp/internal/config"
	"WebApp/internal/forms"
	"WebApp/internal/models"
	"WebApp/internal/render"
	"encoding/json"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository structure contains a variable App which is a pointer to the config file
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
	render.Template(w, r, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip") // m.App.Session is used to access the SessionManager package functions.
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.Template(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Coffee renders the coffee page
func (m *Repository) Coffee(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "coffee.page.html", &models.TemplateData{})
}

// CassavaCake renders the cassava-cake page
func (m *Repository) CassavaCake(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "cassava-cake.page.html", &models.TemplateData{})
}

// Order renders the order page
func (m *Repository) Order(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "order.page.html", &models.TemplateData{})
}

// Confirm renders the confirm page
func (m *Repository) Confirm(w http.ResponseWriter, r *http.Request) {
	var emptyInfo models.ConfirmInfo
	data := make(map[string]interface{})
	data["confirmInfo"] = emptyInfo

	render.Template(w, r, "confirm.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostConfirm handles the post form logic and redirect to order summary
func (m *Repository) PostConfirm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	confirmInfo := models.ConfirmInfo{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	// Validations
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["confirmInfo"] = confirmInfo

		render.Template(w, r, "confirm.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "confirmInfo", confirmInfo) // storing value of the session

	http.Redirect(w, r, "/order-summary", http.StatusSeeOther)
}

// Contact renders the Contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.html", &models.TemplateData{})
}

// OrderSummary displays information on the page for a given order
func (m *Repository) OrderSummary(w http.ResponseWriter, r *http.Request) {
	confirmInfo, ok := m.App.Session.Get(r.Context(), "confirmInfo").(models.ConfirmInfo) // type assertion to models.ConfirmInfo struct
	if !ok {
		log.Println("cannot get information from session")
		m.App.Session.Put(r.Context(), "error", "Can not get order summary from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "confirmInfo")
	data := make(map[string]interface{}) // created a map to access the confirmInfo template data
	data["confirmInfo"] = confirmInfo

	render.Template(w, r, "order-summary.page.html", &models.TemplateData{ // render the order summary and passed the confirmInfo and orderInfo template data
		Data: data,
	})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) OrderAvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
