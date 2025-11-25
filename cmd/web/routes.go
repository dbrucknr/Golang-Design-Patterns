package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) routes() http.Handler {
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

	// Routes Registration
	mux.Get("/", app.ShowHome)

	return mux
}
