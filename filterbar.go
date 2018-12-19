package main

import (
	"github.com/appventure-nush/appventure-website/api/content"
)

type Filter struct {
	Title string
	Tags  map[string]string
}

func GetFilterbar() []Filter {
	return []Filter{
		Filter{
			Title: "Platforms",
			Tags:  content.AppPlatforms,
		},
		Filter{
			Title: "Year",
			Tags:  content.AppYear,
		},
		Filter{
			Title: "Project Type",
			Tags:  content.AppType,
		},
	}
}
