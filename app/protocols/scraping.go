package protocols

import "lucky-sena/domain"

type BetScraper interface {
	Scrap(string) (domain.Bet, error)
}
