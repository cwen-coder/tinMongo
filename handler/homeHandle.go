package handler

import (
	// "github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Home(r render.Render) {
	r.HTML(200, "home/home", nil)
}
