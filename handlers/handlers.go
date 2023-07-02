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

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}
