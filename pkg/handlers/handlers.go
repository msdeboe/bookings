package handlers

import (
	"net/http"

	"github.com/msdeboe/bookings/pkg/config"
	"github.com/msdeboe/bookings/pkg/models"
	"github.com/msdeboe/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the Repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handerls
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.SessionManager.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//Perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello down there"

	remoteIP := m.App.SessionManager.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
