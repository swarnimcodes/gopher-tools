package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/swarnimcodes/gopher-tools/components"
	"github.com/swarnimcodes/gopher-tools/handlers"
)

// embed the static dir into the binary at comp time
//
//go:embed static
var staticFS embed.FS

func main() {
	r := chi.NewRouter()

	// create a sub-fs that starts at directory "static"
	staticSub, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticSub))))

	initialCount := 10
	counterHandler := handlers.NewCounterHandler(initialCount)

	// routes
	r.Get("/", templ.Handler(components.Index("Swarnim", initialCount)).ServeHTTP)
	r.Post("/increment", counterHandler.Increment)
	r.Post("/decrement", counterHandler.Decrement)

	fmt.Println("Server listening at :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
