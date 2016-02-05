package handler

import (
	// "github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Home(r render.Render) {
	r.HTML(200, "server/home", nil)
}

func Status(r render.Render) {
	r.HTML(200, "server/status", nil)
}

func Databases(r render.Render) {
	r.HTML(200, "server/databases", nil)
}

func ProcessList(r render.Render) {
	r.HTML(200, "server/processList", nil)
}

func Command(r render.Render) {
	r.HTML(200, "server/command", nil)
}

func Execute(r render.Render) {
	r.HTML(200, "server/execute", nil)
}

func Replication(r render.Render) {
	r.HTML(200, "server/replication", nil)
}
