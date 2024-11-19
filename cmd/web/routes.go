package main

import (
	"github.com/msdeboe/bookings/pkg/config"
	"github.com/msdeboe/bookings/pkg/handlers"

	"github.com/gorilla/mux"
)

func Routes(app *config.AppConfig) *mux.Router {
	r := mux.NewRouter()
	r.Use(NoSurf)
	r.Use(SessionLoad)
	r.HandleFunc("/", handlers.Repo.Home)
	r.HandleFunc("/about", handlers.Repo.About)
	return r
}
