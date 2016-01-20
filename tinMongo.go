package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"runtime"
	"tinMongo/config"
	"tinMongo/routes"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
	}))
	routes.Route(m)
	port := ":" + config.GetPort()
	m.RunOnAddr(port)
	m.Run()
}
