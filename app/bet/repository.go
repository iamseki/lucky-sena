package betusecases

import "lucky-sena/domain/bet"

type AddBetRepository interface {
	Add(bet.Bet) (bet.BetModel, error)
}

type FindBetRepository interface {
	Find() ([]bet.Bet, error)
}

type FindBetByCodeRepository interface {
	FindBetByCode(code int) (bet.Bet, error)
}