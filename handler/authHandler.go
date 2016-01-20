package handler

import (
	// "github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Login(r render.Render) {
	r.HTML(200, "login/login", nil)
}
