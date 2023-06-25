package main

import (
	"fmt"
	"log"
	"github.com/jsonball/bookings/config"
	"github.com/jsonball/bookings/handlers"
	"github.com/jsonball/bookings/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var app config.AppConfig

func main() {
	portNumber := ":8080"
	// Change this to true in production
	app.InProduction = false
	// Creates and sets cookie
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	// stores cookie in app config type var: app
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Println(err)
	} else {
		app.TemplateCache = tc
	}
	app.UseCache = app.InProduction
	render.SetAppConfig(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println("Listening on port :8080")
	serve := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}
