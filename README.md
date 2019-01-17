
# appventure-website [![Build Status](https://appventure.nushigh.edu.sg:8000/api/badges/appventure-nush/appventure-website/status.svg)](https://appventure.nushigh.edu.sg:8000/appventure-nush/appventure-website)

\[Deployment: [docker-compose.yml](https://github.com/appventure-nush/infrastructure/blob/master/setup-scripts/main-website.yml) | [registry (website)](https://appventure.nushigh.edu.sg/registry/#/appventure-website) | [registry (api)](https://appventure.nushigh.edu.sg/registry/#/appventure-api)\]

> Backstory: The current version of the website uses Jekyll and a submission portal that pushes to Git. However, on our new setup, none of the services should have push access to GitHub, and anyway it is better to move to a proper CMS so that we can add additional sections to the website without Jekyll hackery. Hence, due to my obsession with Go and APIs, here we are.

The AppVenture website is a Go webapp that generates pages out of static [templates](templates/) and JSON data obtained from the API. It also serves the static files in the [assets](assets/) folder.

* `api.go`: A thin wrapper to perform requests to the API
* `router.go`: Each request path (example: /about) is defined here and a handler in `handlers.go` is called
* `handlers.go`: Each page is "handled" (API request and rendering) as defined here
* `helpers.go`: Functions used in templates
* `templatemanager.go`: Templating engine manager, loads and provides functions for each template
* `main.go`: Glue to bind the layers together and provide a commandline interface

## Assets

The stylesheets are written in Sass's SCSS format and complied into a single stylesheet using Grunt. JavaScript is written plain using prototype syntax.

# appventure-api

In the [api](api/) folder is where the resources the API provides are defined. It uses [Ponzu](https://docs.ponzu-cms.org/) to automatically generate an API from the types defined in the [content](api/content/) folder.

It also exposes an admin interface at the path /admin/

# Development

Requirements:

* Go
* Node.js
* Ponzu (`go get github.com/ponzu-cms/ponzu/...`)

```bash
ORG=$(go env GOPATH)/src/github.com/appventure-nush
mkdir -p $ORG
git clone git@github.com:appventure-nush/appventure-website.git $ORG/appventure-website
cd $ORG/appventure-website
# build CSS
yarn && yarn grunt
# build and run website
go get && go build && ./appventure-website -debug
# build and run API
cd api && ponzu build && ponzu run --port 8081
```
