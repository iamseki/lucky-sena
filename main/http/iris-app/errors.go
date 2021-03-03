package irisapp

import "github.com/kataras/iris/v12"

func badRequest(ctx iris.Context) {
	ctx.StatusCode(400)
	ctx.JSON(map[string]string{"tip": "check your body request"})
}

func serverError(ctx iris.Context, err error) {
	ctx.StatusCode(500)
	ctx.JSON(map[string]string{"error": err.Error()})
}
