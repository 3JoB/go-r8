package view

import (
	"fmt"

	"github.com/3JoB/unsafeConvert"
	"github.com/savsgio/atreugo/v11"

	api_model "github.com/3JoB/go-r8/internal/model/api"
)

func NotFound(rc *atreugo.RequestCtx) error {
	var rd api_model.View = api_model.View{
		Code:    404,
		Message: fmt.Sprintf("Path %v not found", unsafeConvert.StringReflect(rc.URI().Path())),
	}
	return rc.JSONResponse(&rd, 404)
}

func Panic(rc *atreugo.RequestCtx, w any) {
	var rd api_model.View = api_model.View{
		Code:    500,
		Message: "internal error",
	}
	rc.JSONResponse(&rd, 500)
}
