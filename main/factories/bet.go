package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/infra/db/mongodb"
)

func NewAddBetUseCase() *betusecases.DbAddBet {
	addBetRepository := mongodb.NewAddBetMongoRepository()
	return betusecases.NewAddBet(addBetRepository)
}

func NewFindBetsUseCase() *betusecases.FindBet {
	findBetRepository := mongodb.NewFindBetMongoRepository()
	return betusecases.NewFindBet(findBetRepository)
}

func NewInsertBetsUseCase() *betusecases.DbInsertBets {
	insertBetsRepository := mongodb.NewInsertBetsMongoRepository()
	return betusecases.NewInsertBets(insertBetsRepository)
}
