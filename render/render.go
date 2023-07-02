package render

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jsonball/bookings/models"
)

var templateCache *template.Template

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	if templateCache != nil {
		err := templateCache.ExecuteTemplate(w, tmpl, td)
		if err != nil {
			log.Println("error executing template:", tmpl, "error message:", err)
		}
	} else {
		t, err := template.ParseFiles("./templates/base.layout.html", "./templates/"+tmpl)
		if err != nil {
			log.Println("error parsing template:", tmpl, "error message:", err)
		} else {
			t.ExecuteTemplate(w, tmpl, td)
		}

	}
}

func CreateTemplateCache() {
	tc, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Println("error creating template cache:", err)
		return
	}
	templateCache = tc
}
func ClearTemplateCache() {
	templateCache = nil
}
