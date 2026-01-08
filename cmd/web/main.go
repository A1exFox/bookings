package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a1exfox/bookings/internal/config"
	"github.com/a1exfox/bookings/internal/handlers"
	"github.com/a1exfox/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = "localhost:8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
