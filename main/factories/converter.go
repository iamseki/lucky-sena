package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"lucky-sena/infra/converter"
	"lucky-sena/infra/db/mongodb"
)

// NewResultsBetToCsvConverter injects all dependencies in any object that implements domain.BetCsvConverter and return it
// collection = results
func NewResultsBetToCsvConverter() domain.BetCsvConverter {
	r := mongodb.NewFindBetMongoRepository("results")
	c := converter.NewBetToCsvConverter("results")

	return betusecases.NewDbBetToCsv(r, c)
}

// NewMadeBetToCsvConverter injects all dependencies in any object that implements domain.BetCsvConverter and return it
// collection = bets
func NewMadeBetToCsvConverter() domain.BetCsvConverter {
	r := mongodb.NewFindBetMongoRepository("bets")
	c := converter.NewBetToCsvConverter("bets")

	return betusecases.NewDbBetToCsv(r, c)
}
