package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type TemplateManager struct {
	tmpls   *template.Template
	funcMap template.FuncMap
	dir     string
}

func (tm *TemplateManager) Render(w http.ResponseWriter, tmpl string, data interface{}) {
	err := tm.tmpls.ExecuteTemplate(w, tmpl, data)
	if err != nil { // should not happen
		log.Print(err.Error())
		http.Error(w, "Whoops, an internal server error occurred!\nContact appventure@nushigh.edu.sg if the problem persists\n\nPage generated "+time.Now().String(), http.StatusInternalServerError)
	}
}

func (tm *TemplateManager) RenderError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	tm.Render(w, "error.html", &Context{
		Title: fmt.Sprintf("%v", code),
		Item:  code,
	})
}

func (tm *TemplateManager) Reload() error {
	tm.tmpls = template.New("templates").Funcs(tm.funcMap)
	return filepath.Walk(tm.dir, func(path string, _ os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			_, err := tm.tmpls.ParseFiles(path)
			return err
		}
		return nil
	})
}

func (tm *TemplateManager) AddFunction(name string, f interface{}) {
	tm.funcMap[name] = f
}

// NewTemplateManager loads templates and allows reloading them during runtime
func NewTemplateManager(dir string, helpers template.FuncMap) *TemplateManager {
	tm := &TemplateManager{
		tmpls:   nil,
		dir:     dir,
		funcMap: helpers,
	}

	err := tm.Reload()
	if err != nil {
		log.Fatal(err)
	}

	return tm
}
