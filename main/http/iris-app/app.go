package irisapp

import (
	"log"
	"lucky-sena/main/http/handlers"
	"os"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"github.com/kataras/iris/v12/middleware/logger"
)

// NewIrisApp returns an instance of iris app
func NewIrisApp() *iris.Application {
	app := iris.New()
	app.Use(cors.Default())

	username := os.Getenv("API_BASIC_USERNAME")
	password := os.Getenv("API_BASIC_PASSWORD")

	auth := basicauth.New(basicauth.Options{
		Allow: basicauth.AllowUsers(map[string]string{
			username: password,
		}),
		Realm: "Authorization Required",
	})

	app.Get("/hc", func(ctx iris.Context) { ctx.JSON(map[string]string{"health": "ok"}) })
	routesV1 := app.Party("/api/v1")
	{
		routesV1.Use(logger.New())
		routesV1.Use(auth)

		routesV1.Post("/bets/generate", newGenerateBetsIrisAdapter(handlers.GenerateBetsHandle))
	}

	if err := app.Build(); err != nil {
		log.Fatalln(err)
	}

	return app
}
