package irisapp

import (
	"lucky-sena/main/http/handlers"
	"os"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"github.com/kataras/iris/v12/middleware/logger"
)

func setupRoutesAndMiddlewares(app *iris.Application) {
	app.Use(cors.Default())
	app.Get("/hc", func(ctx iris.Context) { ctx.JSON(map[string]string{"health": "ok"}) })

	v1 := app.Party("/api/v1")

	v1.Use(logger.New())
	v1.Use(setupAuth())

	v1.Post("/bets/generate", newPostGenerateBetsIrisAdapter(handlers.GenerateBetsHandle))
	v1.Get("/bets/generate", newGetGenerateBetsIrisAdapter(handlers.GenerateBetsHandle))
}

func setupAuth() context.Handler {
	username := os.Getenv("API_BASIC_USERNAME")
	password := os.Getenv("API_BASIC_PASSWORD")

	return basicauth.New(basicauth.Options{
		Allow: basicauth.AllowUsers(map[string]string{
			username: password,
		}),
		Realm: "Authorization Required",
	})
}
