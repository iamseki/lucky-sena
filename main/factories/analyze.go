package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"lucky-sena/infra/db/mongodb"
)

// NewAnalyzeBetUseCase injects all dependencies in any object that implements domain.BetAnalyzer and return it
func NewAnalyzeBetUseCase(collection string) domain.BetAnalyzer {
	analyzeBetRepository := mongodb.NewAnalyzeBetMongoRepository(collection)
	return betusecases.NewAnalyzeBet(analyzeBetRepository)
}
