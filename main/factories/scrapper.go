package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/app/protocols"
	"lucky-sena/infra/scrapper"
)

func NewScrappingLastBetUseCase() protocols.BetScraper {
	c := scrapper.NewCollyScrapper()
	return betusecases.NewLastBetScrapper(c)
}
