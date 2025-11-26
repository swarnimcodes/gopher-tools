package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/swarnimcodes/gopher-toolbox/components"
	"github.com/swarnimcodes/gopher-toolbox/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	initialCount := 10
	counterHandler := handlers.NewCounterHandler(initialCount)

	// routes
	r.Get("/", templ.Handler(components.Index("Swarnim", initialCount)).ServeHTTP)
	r.Post("/increment", counterHandler.Increment)
	r.Post("/decrement", counterHandler.Decrement)

	fmt.Println("Server listening at :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
