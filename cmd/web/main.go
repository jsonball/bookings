package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jsonball/bookings/config"
	"github.com/jsonball/bookings/handlers"
	"github.com/jsonball/bookings/render"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var app config.AppConfig
var once sync.Once

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

	once.Do(render.CreateTemplateCache)
	r := handlers.NewRepo(&app)
	handlers.NewHandlers(r)

	fmt.Println("Listening on port :8080")
	serve := &http.Server{
		Addr:    portNumber,
		Handler: Routes(),
	}
	err := serve.ListenAndServe()
	log.Fatal(err)
}
