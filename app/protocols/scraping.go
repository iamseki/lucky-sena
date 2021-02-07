package protocols

import "lucky-sena/domain/bet"

type BetScraper interface {
	Scrap(string) bet.Bet
}
