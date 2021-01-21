package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/infra/db/mongodb"
)

func NewAnalyzeBetUseCase() *betusecases.AnalyzeBet {
	analyzeBetRepository := mongodb.NewAnalyzeBetMongoRepository()
	return betusecases.NewAnalyzeBet(analyzeBetRepository)
}
