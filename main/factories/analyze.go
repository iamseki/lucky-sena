package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"lucky-sena/infra/db/mongodb"
)

func NewAnalyzeBetUseCase(collection string) domain.BetAnalyzer {
	analyzeBetRepository := mongodb.NewAnalyzeBetMongoRepository(collection)
	return betusecases.NewAnalyzeBet(analyzeBetRepository)
}
