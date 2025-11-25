package main

import (
	"fmt"
	"net/http"
)

func (app *Application) ShowHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
