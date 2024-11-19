package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/msdeboe/bookings/pkg/config"
	"github.com/msdeboe/bookings/pkg/handlers"
	"github.com/msdeboe/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const PortNumber = ":8080"

var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {
	app.InProduction = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.SessionManager = sessionManager

	//load template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	//put template cache in app config and send it to render package
	app.TemplateCache = tc

	//set to false when DEV MODE
	app.UseCache = app.InProduction

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	dir, _ := os.Getwd()
	_, _ = fmt.Printf("Port num: %s\n", PortNumber)
	_, _ = fmt.Printf("Root directory: %s\n", dir)

	srv := &http.Server{Addr: PortNumber, Handler: Routes(&app)}
	err = srv.ListenAndServe()

	log.Fatal(err.Error())
}
