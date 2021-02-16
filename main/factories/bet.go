package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"lucky-sena/infra/db/mongodb"
)

func NewAddBetUseCase() domain.BetSaver {
	addBetRepository := mongodb.NewAddBetMongoRepository()
	return betusecases.NewAddBet(addBetRepository)
}

func NewFindBetsUseCase() domain.BetFinder {
	findBetRepository := mongodb.NewFindBetMongoRepository("results")
	return betusecases.NewFindBet(findBetRepository)
}

func NewInsertBetsUseCase(collection string) domain.BetInserter {
	insertBetsRepository := mongodb.NewInsertBetsMongoRepository(collection)
	return betusecases.NewInsertBets(insertBetsRepository)
}
