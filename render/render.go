package render

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jsonball/bookings/models"
)

var templateCache *template.Template

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	templateCache.ExecuteTemplate(w, tmpl, td)
}

func CreateTemplateCache() {
	tc, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Println(err)
		return
	}
	templateCache = tc
}
