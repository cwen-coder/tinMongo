package handler

import (
	// "github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Home(r render.Render) {
	r.HTML(200, "server/home", map[string]interface{}{})
}

func Status(r render.Render) {
	r.HTML(200, "server/status", map[string]interface{}{})
}

func Databases(r render.Render) {
	r.HTML(200, "server/databases", map[string]interface{}{})
}

func ProcessList(r render.Render) {
	r.HTML(200, "server/processList", map[string]interface{}{})
}

func Command(r render.Render) {
	r.HTML(200, "server/command", map[string]interface{}{})
}

func Execute(r render.Render) {
	r.HTML(200, "server/execute", map[string]interface{}{})
}

func Replication(r render.Render) {
	r.HTML(200, "server/replication", map[string]interface{}{})
}
