package irisapp

import (
	"lucky-sena/main/http/handlers"

	"github.com/kataras/iris/v12"
)

func NewIrisApp() *iris.Application {
	app := iris.New()

	routesV1 := app.Party("/api/v1")
	{
		routesV1.Post("/generate", newGenerateBetsIrisAdapter(handlers.GenerateBetsHandle))
	}

	return app
}
