package content

import (
	"fmt"
	"net/http"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type App struct {
	item.Item

	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Authors          []string `json:"authors"`
	Achievements     string   `json:"achievements"`
	Platforms        []string `json:"platforms"`
	Year             []string `json:"year"`
	Type             string   `json:"type"`
	PlaystorePackage string   `json:"playstore_package"`
	AppstoreUrl      string   `json:"appstore_url"`
	Links            []string `json:"links"`
	Downloads        []string `json:"downloads"`
	Icon             string   `json:"icon"`
	Screenshots      []string `json:"screenshots"`
	Content          string   `json:"content"`
	Flags            []string `json:"flags"`
}

// Helpers

func (a *App) Flagged(name string) bool {
	for _, f := range a.Flags {
		if f == name {
			return true
		}
	}
	return false
}

// MarshalEditor writes a buffer of html to edit a App within the CMS
// and implements editor.Editable
func (a *App) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(a,
		// Take note that the first argument to these Input-like functions
		// is the string version of each App field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", a, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter your app name",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", a, map[string]string{
				"label":       "Short Description",
				"placeholder": "Provide a short description of your app",
			}),
		},
		editor.Field{
			View: editor.Richtext("Content", a, map[string]string{
				"label":       "Content",
				"placeholder": "Describe the app, the motivation behind it and your experience",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Authors", a, map[string]string{
				"label":       "Authors",
				"type":        "text",
				"placeholder": "Enter each author here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Achievements", a, map[string]string{
				"label":       "Achievements",
				"placeholder": "List each achievement",
			}),
		},
		editor.Field{
			View: editor.Checkbox("Platforms", a, map[string]string{
				"label": "Platforms Supported",
			}, map[string]string{
      	"Mobile":  "Mobile",
      	"Desktop": "Desktop",
      	"Web":     "Web",
      }),
		},
		editor.Field{
			View: editor.Checkbox("Year", a, map[string]string{
				"label": "Year in NUS High",
			}, map[string]string{
      	"Year 1": "Year 1",
      	"Year 2": "Year 2",
      	"Year 3": "Year 3",
      	"Year 4": "Year 4",
      	"Year 5": "Year 5",
      	"Year 6": "Year 6",
      }),
		},
		editor.Field{
			View: editor.Select("Type", a, map[string]string{
				"label": "Type",
			}, map[string]string{
      	"CS Module":   "CS Module",
      	"Competition": "Competition",
      	"By Request":  "By Request",
      	"Personal":    "Personal",
      }),
		},
		editor.Field{
			View: editor.Input("PlaystorePackage", a, map[string]string{
				"label":       "Playstore Package",
				"type":        "text",
				"placeholder": "Paste your package path (com.example.app)",
			}),
		},
		editor.Field{
			View: editor.Input("AppstoreUrl", a, map[string]string{
				"label":       "Appstore URL",
				"type":        "text",
				"placeholder": "Paste your Apple App Store URL",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Links", a, map[string]string{
				"label":       "Links",
				"type":        "text",
				"placeholder": "Link your project website or source code",
			}),
		},
		editor.Field{
			View: editor.FileRepeater("Downloads", a, map[string]string{
				"label":       "Downloads",
				"type":        "text",
				"placeholder": "Upload files for users to download",
			}),
		},
		editor.Field{
			View: editor.File("Icon", a, map[string]string{
				"label":       "Icon",
				"placeholder": "Upload your app icon",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Screenshots", a, map[string]string{
				"label": "Screenshots",
			},
				"Screenshot",
				`{{ .hint }} "{{ .description }}"`,
			),
		},
		editor.Field{
			View: editor.Checkbox("Flags", a, map[string]string{
				"label": "Flags",
			}, map[string]string{
				"featured": "Featured",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render App editor view: %s", err.Error())
	}

	return view, nil
}

// Create implements api.Createable
func (a *App) Create(res http.ResponseWriter, req *http.Request) error {
	return nil
}

// Approve implements editor.Mergeable
func (a *App) Approve(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func init() {
	item.Types["App"] = func() interface{} { return new(App) }
}

// String defines how a App is printed. Update it using more descriptive
// fields from the App struct type
func (a *App) String() string {
	return "App: " + a.Name
}
