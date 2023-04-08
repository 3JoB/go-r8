package router

import (
	"github.com/savsgio/atreugo/v11"

	"github.com/3JoB/go-r8/internal/pkg"
)

func Get(server *atreugo.Atreugo) {
	server.ANY("/", pkg.Index)
	// endpoint := server.NewGroupPath("/endpoint")
}
