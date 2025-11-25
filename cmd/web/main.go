package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const PORT = ":4000"

// This type enables shared context between other layers in the application.
type Application struct {
	// Define fields here
}

func main() {
	// Setup Application Instance
	app := Application{}
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
