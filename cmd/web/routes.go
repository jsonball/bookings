package main

import (
	"github.com/jsonball/bookings/config"
	"github.com/jsonball/bookings/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionSaveAndLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
