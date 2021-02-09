package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"lucky-sena/infra/db/mongodb"
)

func NewAnalyzeBetUseCase() *domain.BetAnalyzer {
	analyzeBetRepository := mongodb.NewAnalyzeBetMongoRepository()
	return betusecases.NewAnalyzeBet(analyzeBetRepository)
}
