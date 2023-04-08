package router

import (
	"github.com/3JoB/go-r8/internal/pkg"
	"github.com/savsgio/atreugo/v11"
)

func Get(server *atreugo.Atreugo) {
	server.ANY("/", pkg.Index)
	//endpoint := server.NewGroupPath("/endpoint")
}