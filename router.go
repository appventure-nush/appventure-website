package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router *httprouter.Router
	tm     *TemplateManager
}

// NewRouter initialises routes and links them to handlers
func NewRouter(h *Handlers, assets string) http.Handler {
	router := httprouter.New()

	router.GET("/", h.home)

	router.GET("/about", h.page("about.html", "About"))
	router.GET("/contact", h.page("contact.html", "Contact"))
	router.GET("/request", h.page("request.html", "Request"))
	router.GET("/apps", h.apps)
	router.GET("/apps/:slug", h.app)
	router.GET("/projects", h.projects)
	router.GET("/projects/:slug", h.project)

	// Assets
	router.ServeFiles("/assets/*filepath", http.Dir(assets))

	// System
	router.GET("/heathz", h.healthz)

	// Debug
	if debug {
		router.GET("/_tm_reload", h.tmReload)
	}

	return router
}
