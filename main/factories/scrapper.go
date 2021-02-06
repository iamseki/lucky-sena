package factories

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/infra/scrapper"
)

func NewScrappingLastBetUseCase() betusecases.BetScraper {
	c := scrapper.NewCollyScrapper()
	return betusecases.NewLastBetScrapper(c)
}
