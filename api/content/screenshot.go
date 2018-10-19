package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Screenshot struct {
	item.Item

	Image       string `json:"image"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// MarshalEditor writes a buffer of html to edit a Screenshot within the CMS
// and implements editor.Editable
func (s *Screenshot) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(s,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Screenshot field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.File("Image", s, map[string]string{
				"label":       "Image",
				"placeholder": "Upload the screenshot",
			}),
		},
		editor.Field{
			View: editor.Select("Type", s, map[string]string{
				"label": "Type",
			}, map[string]string{
				"mobile-9-16":   "Mobile (9:16)",
				"mobile-16-9":   "Mobile Landscape (16:9)",
				"desktop-16-10": "Desktop (16:10)",
			}),
		},
		editor.Field{
			View: editor.Input("Description", s, map[string]string{
				"label":       "Description",
				"type":        "text",
				"placeholder": "Describe what is seen",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Screenshot editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Screenshot"] = func() interface{} { return new(Screenshot) }
}

// String defines how a Screenshot is printed. Update it using more descriptive
// fields from the Screenshot struct type
func (s *Screenshot) String() string {
	return fmt.Sprintf("Screenshot: %s", s.UUID)
}
