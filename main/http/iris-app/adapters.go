package irisapp

import (
	"lucky-sena/domain"
	"lucky-sena/main/factories"
	"lucky-sena/main/http/handlers"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func newPostGenerateBetsIrisAdapter(generateBetsFn handlers.GenerateBetsHandler) context.Handler {
	return func(ctx iris.Context) {
		request := &handlers.RequestBetsHandle{}
		err := ctx.ReadJSON(&request)
		if err := handlers.ValidateRequestBets(request); err != nil {
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

func newGetGenerateBetsIrisAdapter(generateBetsFn handlers.GenerateBetsHandler) context.Handler {
	return func(ctx iris.Context) {
		bets, err := strconv.Atoi(ctx.URLParamDefault("bets", "1"))
		if err != nil {
			badRequest(ctx)
			return
		}

		generatedBets := generateBetsFn(bets, []int{})
		response := &handlers.ResponseBetsHandle{Bets: generatedBets}
		ctx.JSON(response)
	}
}
