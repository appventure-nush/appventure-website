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
		"html": func(a string) template.HTML {
			// WARNING: we trust that site admins will not inject malicious code into the website
			// If such trust can not be affirmed, add XSS filters here
			return template.HTML(a)
		},
		"even": func(i int) bool {
			return i%2 == 0
		},
		"size": func(size, url string) string {
			if debug {
				return apiHost + url + "?size=" + size
			}
			return "/" + size + strings.Replace(url, "/api/uploads", "", -1)
		},
		"filterbar": GetFilterbar,
	}
}
