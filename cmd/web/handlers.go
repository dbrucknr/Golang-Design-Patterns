package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "home.page.gohtml", nil)
}

func (app *Application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	app.RenderTemplate(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}
