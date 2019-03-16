package content

import (
	"fmt"
	"net/http"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Project struct {
	item.Item

	Name    string   `json:"name"`
	Authors []string `json:"authors"`
	Summary string   `json:"Summary"`
	Content string   `json:"content"`
}

// MarshalEditor writes a buffer of html to edit a Project within the CMS
// and implements editor.Editable
func (p *Project) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Project field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", p, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the project title",
			}),
		},
		editor.Field{
			View: editor.InputRepeater("Authors", p, map[string]string{
				"label":       "Authors",
				"type":        "text",
				"placeholder": "Enter each author here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Summary", p, map[string]string{
				"label":       "Summary",
				"placeholder": "Enter your project summary here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Content", p, map[string]string{
				"label":       "Content",
				"placeholder": "Enter your project post here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Project editor view: %s", err.Error())
	}

	return view, nil
}

// Create implements api.Createable
func (p *Project) Create(res http.ResponseWriter, req *http.Request) error {
	return nil
}

// Approve implements editor.Mergeable
func (p *Project) Approve(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func init() {
	item.Types["Project"] = func() interface{} { return new(Project) }
}

// String defines how a Project is printed. Update it using more descriptive
// fields from the Project struct type
func (p *Project) String() string {
	return "Project: " + p.Name
}
