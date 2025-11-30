package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) Routes() http.Handler {
	mux := chi.NewRouter()

	// Middleware Registration
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.NoCache)
	mux.Use(middleware.Compress(5, "application/json", "text/plain", "text/html"))
	mux.Use(middleware.RedirectSlashes)
	mux.Use(middleware.Throttle(100))
	mux.Use(middleware.Timeout(60 * time.Second)) // 60 Seconds Timeout

	// Static Assets Server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Display Test Page
	mux.Get("/test-patterns", app.TestPatterns)

	// API Routes
	mux.Get("/api/dog-from-factory", app.CreateDogFromFactory)
	mux.Get("/api/cat-from-factory", app.CreateCatFromFactory)
	mux.Get("/api/dog-from-abstract-factory", app.CreateDogFromAbstractFactory)
	mux.Get("/api/cat-from-abstract-factory", app.CreateCatFromAbstractFactory)
	mux.Get("/api/dog-breeds", app.GetAllDogBreedsJson)
	// mux.Get("/api/cat-breeds", app.GetAllCatBreedsJson)
	// Builder Route
	mux.Get("/api/dog-from-builder", app.CreateDogFromBuilder)
	mux.Get("/api/cat-from-builder", app.CreateCatFromBuilder)

	// Routes Registration
	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)

	return mux
}
