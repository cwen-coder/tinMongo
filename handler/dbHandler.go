package handler

import (
	// "github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func DbHome(r render.Render) {
	r.HTML(200, "db/home", nil)
}
