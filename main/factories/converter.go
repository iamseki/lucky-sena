package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain/bet"
	"lucky-sena/infra/converter"
	"lucky-sena/infra/db/mongodb"
)

func NewBetToCsvConverter() bet.BetCsvConverter {
	r := mongodb.NewFindBetMongoRepository()
	c := converter.NewBetToCsvConverter()

	return betusecases.NewDbBetToCsv(r, c)
}
