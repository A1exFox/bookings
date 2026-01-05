package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a1exfox/go-course/pkg/config"
	"github.com/a1exfox/go-course/pkg/handlers"
	"github.com/a1exfox/go-course/pkg/render"
)

const portNumber = "localhost:8080"

func main() {

    var app config.AppConfig

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
        Addr: portNumber,
        Handler: routes(&app),
    }
    err = srv.ListenAndServe()
    log.Fatal(err)
}
