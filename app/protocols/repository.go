package protocols

import (
	"lucky-sena/domain"
)

type AddBetRepository interface {
	Add(domain.Bet) (domain.BetModel, error)
}

type InsertBetsRepository interface {
	InsertMany([]domain.Bet) error
}

type FindBetRepository interface {
	Find() ([]domain.Bet, error)
}

type FindBetByCodeRepository interface {
	FindBetByCode(code int) (domain.Bet, error)
}

type FindBetsByNumbersRepository interface {
	FindBetsByNumbers(numbers []int) ([]domain.Bet, error)
}

type AnalyzeBetRepository interface {
	FindBetRepository
	FindBetByCodeRepository
	FindBetsByNumbersRepository
}
