package handlers

import (
	"net/http"

	"github.com/jsonball/bookings/config"
	"github.com/jsonball/bookings/models"
	"github.com/jsonball/bookings/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringData := make(map[string]string)
	stringData["greeting"] = "hello"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringData["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringData})
}
