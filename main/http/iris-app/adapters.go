package irisapp

import (
	"lucky-sena/domain"
	"lucky-sena/main/factories"
	"lucky-sena/main/http/handlers"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func newGenerateBetsIrisAdapter(generateBetsFn handlers.GenerateBetsHandler) context.Handler {
	return func(ctx iris.Context) {
		request := &handlers.RequestBetsHandle{}
		ctx.ReadJSON(&request)
		if request.Bets == 0 || request.ExcludedBets == nil {
			badRequest(ctx)
			return
		}

		analyzeBetUseCase := factories.NewAnalyzeBetUseCase("results")
		code, err := analyzeBetUseCase.NextBetCode()
		if err != nil {
			serverError(ctx, err)
			return
		}

		addBetUseCase := factories.NewAddBetUseCase()
		generatedBets := generateBetsFn(request.Bets, request.ExcludedBets)
		for _, b := range generatedBets {
			addBetUseCase.AddBet(domain.Bet{Numbers: b.Numbers, Code: code, Date: time.Now()})
		}

		response := &handlers.ResponseBetsHandle{Bets: generatedBets}
		ctx.JSON(response)
	}
}
