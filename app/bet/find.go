package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain"
)

// FindBet represents a struct that implements domain.BetFinder interface
type FindBet struct {
	repository protocols.FindBetRepository
}

// NewFindBet returns an instance of FinBet struct
func NewFindBet(repository protocols.FindBetRepository) *FindBet {
	return &FindBet{
		repository,
	}
}

// Find query for all bets in the database
func (f *FindBet) Find() ([]domain.Bet, error) {
	bets, err := f.repository.Find()
	if err != nil {
		return nil, err
	}
	return bets, err
}
