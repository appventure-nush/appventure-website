package main

import (
	"html/template"
	"regexp"
	"strings"
	"time"
)

func slugToImage(size, slug string) string {
	if debug {
		return apiHost + slug + "?size=" + size
	}
	// TODO: use signed URLs
	return "/img/" + size + strings.Replace(slug, "/api/uploads", "", -1)
}

var imgRegexp = regexp.MustCompile(` src=".*?"`)
var pRegexp = regexp.MustCompile(`(?ms:^\s*<p( .*|)>.*<\/p>\s*$)`)

var singaporeTime, _ = time.LoadLocation("Asia/Singapore")

// NewHelpers returns a list of helpers used in the templates
func NewHelpers() template.FuncMap {
	return template.FuncMap{
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
		"even": func(i int) bool {
			return i%2 == 0
		},
		"unslug": func(a string) string {
			return strings.Join(strings.Split(a, "-")[1:], "-")
		},
		"size": slugToImage,
		"fancyextension": func(s string) string {
			dots := strings.Split(s, ".")
			return strings.ToUpper(dots[len(dots)-1])
		},
		"html": func(a string) template.HTML {
			// Transform <img> into proper links
			// WARNING: classes MUST not be set in the visual editor
			b := imgRegexp.ReplaceAllStringFunc(a, func(match string) string {
				return " src=\"" + slugToImage("520x", match[6:len(match)-1]) + "\" class=\"picture\""
			})
			// Add paragraph tags if necessary
			if !pRegexp.MatchString(b) {
				b = `<p>` + b + `</p>`
			}
			// WARNING: we trust that site admins will not inject malicious code into the website
			// If such trust can not be affirmed, add XSS filters here
			return template.HTML(b)
		},
		"date": func(i int64) string {
			t := time.Unix(i/1000, 0).In(singaporeTime)
			return t.Format("2 Jan 2006")
		},
		"year": func(i int64) string {
			t := time.Unix(i/1000, 0).In(singaporeTime)
			return t.Format("2006")
		},
		"filterbar": getFilterbar,
	}
}
