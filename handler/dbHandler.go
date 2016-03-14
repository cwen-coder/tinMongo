package handler

import (
	// "github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func DbHome(r render.Render) {
	r.HTML(200, "db/home", map[string]interface{}{})
}

func DbNewCollection(r render.Render) {
	r.HTML(200, "db/newCollection", map[string]interface{}{})
}

func DbTransfer(r render.Render) {
	r.HTML(200, "db/dbTransfer", map[string]interface{}{})
}

func DbExport(r render.Render) {
	r.HTML(200, "db/dbExport", map[string]interface{}{})
}

func DbImport(r render.Render) {
	r.HTML(200, "db/dbImport", map[string]interface{}{})
}

func DbUsers(r render.Render) {
	r.HTML(200, "db/dbUsers", map[string]interface{}{})
}
