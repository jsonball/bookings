package main

import (
	"net/http"

	"github.com/jsonball/bookings/handlers"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {

	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionSaveAndLoad)

	fileServer := http.FileServer(http.Dir("public"))
	mux.Handle("/public/*", http.StripPrefix("/public/", fileServer))

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
