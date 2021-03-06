package main

import (
	"html/template"

	"github.com/ghophp/buildbot-dashing/config"
	"github.com/ghophp/buildbot-dashing/container"
	"github.com/ghophp/buildbot-dashing/handler"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func NewRouter(c *container.ContainerBag) *martini.ClassicMartini {
	var (
		indexHandler    = handler.NewIndexHandler(c)
		buildersHandler = handler.NewBuildersHandler(c)
		wsHandler       = handler.NewWsHandler(c)
	)

	router := martini.Classic()
	router.Use(martini.Static("static/assets"))
	router.Use(render.Renderer(render.Options{
		Directory:  "static/templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
		Funcs: []template.FuncMap{
			{
				"genericSize": func() string {
					return c.GenericSize
				},
			},
		},
	}))

	router.Get("/", indexHandler.ServeHTTP)
	router.Get("/ws", wsHandler.ServeHTTP)
	router.Get("/builders", buildersHandler.ServeHTTP)

	return router
}

func main() {
	NewRouter(container.NewContainerBag(config.NewConfig())).Run()
}
