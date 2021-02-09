package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"lucky-sena/infra/converter"
	"lucky-sena/infra/db/mongodb"
)

func NewBetToCsvConverter() domain.BetCsvConverter {
	r := mongodb.NewFindBetMongoRepository()
	c := converter.NewBetToCsvConverter()

	return betusecases.NewDbBetToCsv(r, c)
}
