package pkg

import (
	"github.com/savsgio/atreugo/v11"

	api_model "github.com/3JoB/go-r8/internal/model/api"
)

func Index(rc *atreugo.RequestCtx) error {
	var rd api_model.View = api_model.View{
		Code:    200,
		Message: "r8 is ok",
	}
	return rc.JSONResponse(&rd, 200)
}
