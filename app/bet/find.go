package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain"
)

type FindBet struct {
	repository protocols.FindBetRepository
}

func NewFindBet(repository protocols.FindBetRepository) *FindBet {
	return &FindBet{
		repository,
	}
}

func (f *FindBet) Find() ([]domain.Bet, error) {
	bets, err := f.repository.Find()
	if err != nil {
		return nil, err
	}
	return bets, err
}
