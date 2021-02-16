package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"lucky-sena/infra/converter"
	"lucky-sena/infra/db/mongodb"
)

func NewResultsBetToCsvConverter() domain.BetCsvConverter {
	r := mongodb.NewFindBetMongoRepository("results")
	c := converter.NewBetToCsvConverter()

	return betusecases.NewDbBetToCsv(r, c)
}
