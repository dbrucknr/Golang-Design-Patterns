package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const PORT = ":4000"

// This type enables shared context between other layers in the application.
type Application struct {
	templateMap map[string]*template.Template
	config      ApplicationConfig
}
type ApplicationConfig struct {
	useCache bool
}

func main() {
	// Setup Application Instance
	app := Application{}
	// Command Line Flags
	flag.BoolVar(&app.config.useCache, "use-cache", false, "Enable caching")
	flag.Parse()
	// Server Configuration
	srv := &http.Server{
		Addr:              PORT,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	fmt.Println("Starting server on port", PORT)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
