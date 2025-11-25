package main

import (
	"fmt"
	"net/http"

	"github.com/dbrucknr/go-design-patterns/pets"
	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *Application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "home.page.gohtml", nil)
}

func (app *Application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	app.RenderTemplate(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *Application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *Application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *Application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.RenderTemplate(w, "test.page.gohtml", nil)
}
