package routes

import (
	"github.com/go-martini/martini"
	"tinMongo/handler"
)

func Route(m *martini.ClassicMartini) {
	m.Get("/login", handler.Login)

	m.Get("/home", handler.Home)
	m.Get("/status", handler.Status)
	m.Get("/databases", handler.Databases)
	m.Get("/processList", handler.ProcessList)
	m.Get("/command", handler.Command)
	m.Get("/execute", handler.Execute)
}
