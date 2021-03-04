package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"lucky-sena/infra/db/mongodb"
)

// NewAddBetUseCase injects all dependencies in any object that implements domain.BetSaver and return it
func NewAddBetUseCase() domain.BetSaver {
	addBetRepository := mongodb.NewAddBetMongoRepository()
	return betusecases.NewAddBet(addBetRepository)
}

// NewFindBetsUseCase injects all dependencies in any object that implements domain.BetFinder and return it
func NewFindBetsUseCase() domain.BetFinder {
	findBetRepository := mongodb.NewFindBetMongoRepository("results")
	return betusecases.NewFindBet(findBetRepository)
}

// NewInsertBetsUseCase injects all dependencies in any object that implements domain.BetInserter  and return it
func NewInsertBetsUseCase(collection string) domain.BetInserter {
	insertBetsRepository := mongodb.NewInsertBetsMongoRepository(collection)
	return betusecases.NewInsertBets(insertBetsRepository)
}
