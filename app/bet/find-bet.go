package betusecases

import "lucky-sena/domain/bet"

type FindBet struct {
	repository FindBetRepository
}

func NewFindBet(repository FindBetRepository) *FindBet {
	return &FindBet{
		repository,
	}
}

func (f *FindBet) Find() ([]bet.Bet, error) {
	bets, err := f.repository.Find()
	if err != nil {
		return nil, err
	}
	return bets, err
}
