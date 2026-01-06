package main

import (
	"net/http"

	"github.com/a1exfox/bookings/pkg/config"
	"github.com/a1exfox/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
    mux := chi.NewRouter()

    mux.Use(middleware.Recoverer)
    mux.Use(NoSurf)
    mux.Use(SessionLoad)

    fileServer := http.FileServer(http.Dir("./static/"))
    mux.Handle("/static/*", http.StripPrefix("/static", fileServer))


    mux.Get("/", handlers.Repo.Home)
    mux.Get("/about", handlers.Repo.About)
    return mux
}