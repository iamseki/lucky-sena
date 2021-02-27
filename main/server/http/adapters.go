package main

import (
	"lucky-sena/infra/generator"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func newGenerateBetsIrisAdapter(generateBetsFn generateBetsHandler) context.Handler {
	return func(ctx iris.Context) {
		var request struct {
			Bets         int   `json:"bets"`
			ExcludedBets []int `json:"excludeds"`
		}
		ctx.ReadJSON(&request)
		if request.Bets == 0 || request.ExcludedBets == nil {
			ctx.StatusCode(400)
			return
		}
		var response struct {
			Bets []generator.GenaretedBet
		}
		generatedBets := generateBetsFn(request.Bets, request.ExcludedBets)
		response.Bets = generatedBets
		ctx.JSON(response)
	}
}
