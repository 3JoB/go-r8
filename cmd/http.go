package cmd

import (
	"github.com/savsgio/atreugo/v11"

	"github.com/3JoB/go-r8/internal/config"
	"github.com/3JoB/go-r8/internal/pkg/view"
	"github.com/3JoB/go-r8/internal/router"
)

var conf = config.F()

func ListenServer() {
	server := atreugo.New(atreugo.Config{
		Addr:         conf.String("server.addr"),
		Name:         "r8-core/alpha",
		NotFoundView: view.NotFound,
		PanicView:    view.Panic,
	})

	router.Get(server)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
