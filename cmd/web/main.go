package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/dbrucknr/go-design-patterns/models"
)

const PORT = ":4000"

// This type enables shared context between other layers in the application.
type Application struct {
	templateMap map[string]*template.Template
	config      ApplicationConfig
	Models      models.Models
}
type ApplicationConfig struct {
	useCache bool
	dsn      string
}

func main() {
	// Setup Application Instance
	app := Application{
		templateMap: make(map[string]*template.Template),
	}
	// Command Line Flags - automatically set useCache to false (for development)
	flag.BoolVar(&app.config.useCache, "use-cache", false, "Enable caching")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "Database connection string")
	flag.Parse()

	// Database
	db, err := initMySqlDb(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}
	// Set the connection for the models (sets a package level variable)
	app.Models = *models.New(db)

	// Server Configuration
	srv := &http.Server{
		Addr:              PORT,
		Handler:           app.Routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	fmt.Println("Starting server on port", PORT)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
