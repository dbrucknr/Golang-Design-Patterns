package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":4000"

// This type enables shared context between other layers in the application.
type Application struct {
	// Define fields here
}

func main() {
	// Setup Application Instance
	// app := Application{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	fmt.Println("Starting server on port", PORT)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Panic(err)
	}
}
