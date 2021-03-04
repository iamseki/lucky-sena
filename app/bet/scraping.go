package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain"
)

type LastGameScrapper struct {
	scraper protocols.BetScraper
}

func (s *LastGameScrapper) Scrap(url string) (domain.Bet, error) {
	return s.scraper.Scrap(url)
}

func NewLastBetScrapper(b protocols.BetScraper) *LastGameScrapper {
	return &LastGameScrapper{
		scraper: b,
	}
}
