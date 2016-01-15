package routers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func init() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
	}))

	m.Run()
}
