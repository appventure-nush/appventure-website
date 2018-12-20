package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/appventure-nush/appventure-website/api/content"
	"github.com/julienschmidt/httprouter"
)

type Handlers struct {
	api *API
	tm  *TemplateManager
}

type Context struct {
	Title string
	Jumbo bool
	Items interface{}
	Item  interface{}
}

func (h *Handlers) page(tmpl string, page string) httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		h.tm.Render(w, tmpl, Context{
			Title: page,
		})
	}
}

func (h *Handlers) home(w http.ResponseWriter, rq *http.Request, ps httprouter.Params) {
	apps, err := h.api.FeaturedApps()
	if err != nil {
		log.Println(err)
		h.tm.RenderError(w, http.StatusInternalServerError)
		return
	}
	h.tm.Render(w, "index.html", Context{
		Items: apps,
		Jumbo: true,
	})
}

func (h *Handlers) apps(w http.ResponseWriter, rq *http.Request, ps httprouter.Params) {
	apps, err := h.api.Apps()
	if err != nil {
		log.Println(err)
		h.tm.RenderError(w, http.StatusInternalServerError)
		return
	}
	h.tm.Render(w, "apps.html", Context{
		Title: "Apps",
		Items: apps,
	})
}

func (h *Handlers) app(w http.ResponseWriter, rq *http.Request, ps httprouter.Params) {
	app, err := h.api.App("app-" + ps.ByName("slug"))
	if err == ErrorNotFound {
		log.Println(err)
		h.tm.RenderError(w, http.StatusNotFound)
		return
	} else if err != nil {
		log.Println(err)
		h.tm.RenderError(w, http.StatusInternalServerError)
		return
	}
	screenshots := make([]content.Screenshot, 0)
	for _, s := range app.Screenshots {
		screenshot, err := h.api.ScreenshotByReference(s)
		if err != nil {
			log.Println(err)
		}
		screenshots = append(screenshots, screenshot)
	}
	h.tm.Render(w, "app.html", Context{
		Item:  app,
		Items: screenshots,
	})
}

func (h *Handlers) projects(w http.ResponseWriter, rq *http.Request, ps httprouter.Params) {
	h.tm.Render(w, "projects.html", Context{
		Title: "Projects",
		Items: nil,
	})
}

func (h *Handlers) project(w http.ResponseWriter, rq *http.Request, ps httprouter.Params) {
	h.tm.Render(w, "project.html", Context{
		Title: "",
		Item:  nil,
	})
}

// System
func (h *Handlers) healthz(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "OK")
}

// Debug
func (h *Handlers) tmReload(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := h.tm.Reload()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, "OK")
}

// NewHandlers perform template rendering and fetching of data
func NewHandlers(api *API, tm *TemplateManager) *Handlers {
	return &Handlers{
		api: api,
		tm:  tm,
	}
}
