package irisapp

import (
	"log"

	"github.com/kataras/iris/v12"
)

// New returns an instance of iris app
func New() *iris.Application {
	app := iris.New()
	setupRoutesAndMiddlewares(app)

	if err := app.Build(); err != nil {
		log.Fatalln(err)
	}

	return app
}
