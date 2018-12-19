package main

import (
	"html/template"
	"strings"
)

// NewHelpers returns a list of helpers used in the templates
func NewHelpers() template.FuncMap {
	return template.FuncMap{
		"plusone": func(i int) int {
			return i + 1
		},
		"andify": func(items []string) string {
			str := ""
			for i, item := range items {
				if i == 0 { // first
					str += item
					continue
				}
				if i == len(items)-1 { // last
					str += " and " + item
					continue
				}
				str += ", " + item
			}
			return str
		},
		"unslug": func(a string) string {
			return strings.Join(strings.Split(a, "-")[1:], "-")
		},
		"donotescape": func(a string) template.HTML {
			return template.HTML(a)
		},
		"size": func(size, url string) string {
			return "http://localhost:8080" + url
		},
		"filterbar": GetFilterbar,
	}
}
