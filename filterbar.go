package main

import (
	"appventure-website/api/content"
)

type filter struct {
	Title string
	Tags  map[string]string
}

func getFilterbar() []filter {
	return []filter{
		filter{
			Title: "Platforms",
			Tags:  content.AppPlatforms,
		},
		filter{
			Title: "Year",
			Tags:  content.AppYear,
		},
		filter{
			Title: "Project Type",
			Tags:  content.AppType,
		},
	}
}
