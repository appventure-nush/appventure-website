package main

import (
	"html/template"
	"regexp"
	"strings"
)

func SlugToImage(size, slug string) string {
	if debug {
		return apiHost + slug + "?size=" + size
	}
	return "/img/" + size + strings.Replace(slug, "/api/uploads", "", -1)
}

var imgRegexp = regexp.MustCompile(` src=".*?"`)

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
			// Transform <img> into proper links
			// WARNING: classes MUST not be set in the visual editor
			b := imgRegexp.ReplaceAllStringFunc(a, func(match string) string {
				return " src=\"" + SlugToImage("520x", match[6:len(match)-1]) + "\" class=\"picture\""
			})
			// WARNING: we trust that site admins will not inject malicious code into the website
			// If such trust can not be affirmed, add XSS filters here
			return template.HTML(b)
		},
		"even": func(i int) bool {
			return i%2 == 0
		},
		"size":      SlugToImage,
		"filterbar": GetFilterbar,
	}
}
