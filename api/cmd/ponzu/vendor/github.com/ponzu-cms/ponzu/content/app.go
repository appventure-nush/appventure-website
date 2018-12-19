package content

import (
	"fmt"

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
	Year             []int    `json:"year"`
	Type             string   `json:"type"`
	PlaystorePackage string   `json:"playstore_package"`
	AppstoreUrl      string   `json:"appstore_url"`
	Links            []string `json:"links"`
	Downloads        []string `json:"downloads"`
	Icon             string   `json:"icon"`
	Screenshots      []string `json:"screenshots"`
	Content          string   `json:"content"`
}

var AppPlatforms = map[string]string{
	"mobile":  "Mobile",
	"desktop": "Desktop",
	"web":     "Web",
}

var AppYear = map[string]string{
	"1": "Year 1",
	"2": "Year 2",
	"3": "Year 3",
	"4": "Year 4",
	"5": "Year 5",
	"6": "Year 6",
}
var AppType = map[string]string{
	"cs module":   "CS Module",
	"competition": "Competition",
	"by request":  "By Request",
	"personal":    "Personal",
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
			}, AppPlatforms),
		},
		editor.Field{
			View: editor.SelectRepeater("Year", a, map[string]string{
				"label": "Year in NUS High",
			}, AppYear),
		},
		editor.Field{
			View: editor.Select("Type", a, map[string]string{
				"label": "Type",
			}, AppType),
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
			View: editor.Richtext("Content", a, map[string]string{
				"label":       "Content",
				"placeholder": "Describe the app, the motivation behind it and your experience",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render App editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["App"] = func() interface{} { return new(App) }
}

// String defines how a App is printed. Update it using more descriptive
// fields from the App struct type
func (a *App) String() string {
	return "App: " + a.Name
}
