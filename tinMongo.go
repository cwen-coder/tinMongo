package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"runtime"
	"tinMongo/config"
	"tinMongo/routes"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tpl", ".html"},
		Charset:    "UTF-8",
		Funcs: []template.FuncMap{
			{
				"set": func(renderArgs map[string]interface{}, key string, value interface{}) template.JS {
					renderArgs[key] = value
					return template.JS("")
				},
				"equal": func(args ...interface{}) bool {
					return args[0] == args[1]
				},
			},
		},
	}))
	routes.Route(m)
	port := ":" + config.GetPort()
	m.RunOnAddr(port)
	m.Run()
}
