package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain"
)

// LastGameScrapper represents a struct that implements app/protocols.BetScraper interface
type LastGameScrapper struct {
	scraper protocols.BetScraper
}

// Scrap try to fetch the last bet won in the url
func (s *LastGameScrapper) Scrap(url string) (domain.Bet, error) {
	return s.scraper.Scrap(url)
}

// NewLastBetScrapper returns an isntance of LastGameScrapper struct
func NewLastBetScrapper(b protocols.BetScraper) *LastGameScrapper {
	return &LastGameScrapper{
		scraper: b,
	}
}
