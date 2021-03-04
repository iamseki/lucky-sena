package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/app/protocols"
	"lucky-sena/infra/scrapper"
)

// NewScrappingLastBetUseCase injects all dependencies in any object that implements protocols.BetScraper and return it
func NewScrappingLastBetUseCase() protocols.BetScraper {
	c := scrapper.NewCollyScrapper()
	return betusecases.NewLastBetScrapper(c)
}
