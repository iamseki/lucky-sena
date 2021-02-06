package betusecases

import "lucky-sena/domain/bet"

type BetScraper interface {
	Scrap(string) bet.Bet
}

type LastGameScrapper struct {
	Scraper BetScraper
}

func (s *LastGameScrapper) Scrap(url string) bet.Bet {
	return s.Scraper.Scrap(url)
}

func NewLastBetScrapper(b BetScraper) *LastGameScrapper {
	return &LastGameScrapper{
		Scraper: b,
	}
}
