package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain/bet"
)

type LastGameScrapper struct {
	Scraper protocols.BetScraper
}

func (s *LastGameScrapper) Scrap(url string) bet.Bet {
	return s.Scraper.Scrap(url)
}

func NewLastBetScrapper(b protocols.BetScraper) *LastGameScrapper {
	return &LastGameScrapper{
		Scraper: b,
	}
}
