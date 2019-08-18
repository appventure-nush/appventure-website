package content

import (
  "fmt"
  "net/http"

  "github.com/bosssauce/reference"

  "github.com/ponzu-cms/ponzu/management/editor"
  "github.com/ponzu-cms/ponzu/system/item"
)

type Project struct {
  item.Item

  Name             string   `json:"name"`
  Authors          []string `json:"authors"`
  Year             []int    `json:"year"`
  Platforms        []string `json:"platforms"`
  Summary          string   `json:"summary"`
  Displayimage     string   `json:"displayimage"`
  Screenshots      []string `json:"screenshots"`
  Content          string   `json:"content"`
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
      View: editor.Checkbox("Year", p, map[string]string{
        "label": "Year in NUS High",
      }, map[string]string{
        "1": "Year 1",
        "2": "Year 2",
        "3": "Year 3",
        "4": "Year 4",
        "5": "Year 5",
        "6": "Year 6",
      }),
    },
    editor.Field{
      View: editor.Richtext("SummaryASD", p, map[string]string{
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
    editor.Field{
      View: editor.Checkbox("Platforms", p, map[string]string{
        "label": "Platforms Supported",
      }, map[string]string{
        "mobile":  "Mobile",
        "desktop": "Desktop",
        "web":     "Web",
      }),
    },
    editor.Field{
			View: editor.File("Displayimage", p, map[string]string{
				"label":       "Display Image",
				"placeholder": "Upload display image",
			}),
		},
    editor.Field{
      View: reference.SelectRepeater("Screenshots", p, map[string]string{
        "label": "Screenshots",
      },
        "Screenshot",
        `{{ .hint }} "{{ .description }}"`,
      ),
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
