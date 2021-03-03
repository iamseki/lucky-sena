package irisapp

import (
	"lucky-sena/domain"
	"lucky-sena/infra/generator"
	"lucky-sena/main/factories"
	"lucky-sena/main/http/handlers"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func newGenerateBetsIrisAdapter(generateBetsFn handlers.GenerateBetsHandler) context.Handler {
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

		analyzeBetUseCase := factories.NewAnalyzeBetUseCase("results")
		code, err := analyzeBetUseCase.NextBetCode()
		if err != nil {
			ctx.StatusCode(500)
			ctx.JSON(map[string]string{"error": err.Error()})
			return
		}
		addBetUseCase := factories.NewAddBetUseCase()
		for _, b := range generatedBets {
			addBetUseCase.AddBet(domain.Bet{Numbers: b.Numbers, Code: code, Date: time.Now()})
		}
		ctx.JSON(response)
	}
}
