package irisapp

import (
	"lucky-sena/main/http/handlers"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
)

// NewIrisApp returns an instance of iris app
func NewIrisApp() *iris.Application {
	app := iris.New()

	username := os.Getenv("API_BASIC_USERNAME")
	password := os.Getenv("API_BASIC_PASSWORD")

	auth := basicauth.New(basicauth.Config{
		Users: map[string]string{
			username: password,
		},
		Realm: "Authorization Required",
	})

	routesV1 := app.Party("/api/v1")
	{
		routesV1.Use(auth)
		routesV1.Post("/generate", newGenerateBetsIrisAdapter(handlers.GenerateBetsHandle))
	}

	return app
}
