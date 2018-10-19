package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen    string
	templates string
	assets    string
	apiHost   string
	debug     bool
)

func main() {
	// Parse flags
	flag.StringVar(&listen, "listen", ":8080", "listen address")
	flag.StringVar(&templates, "templates", "templates", "path to templates")
	flag.StringVar(&assets, "assets", "assets", "path to assets")
	flag.StringVar(&apiHost, "api", "http://localhost:8081", "full ponzu API host")
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	if debug {
		log.Println("WARNING: Debug mode enabled")
	}

	// API
	api := NewAPI(apiHost)

	// Helpers
	helpers := NewHelpers()

	// Create template manager
	tm := NewTemplateManager(templates, helpers)

	// Handlers
	handlers := NewHandlers(api, tm)

	// Create router
	router := NewRouter(handlers, assets)

	// Listen
	log.Println("Listening on", listen)
	log.Fatal(http.ListenAndServe(listen, router))
}
