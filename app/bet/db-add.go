package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain"
)

type DbAddBet struct {
	repository protocols.AddBetRepository
}

func NewAddBet(repository protocols.AddBetRepository) *DbAddBet {
	return &DbAddBet{repository}
}

func (db *DbAddBet) AddBet(bet domain.Bet) (domain.BetModel, error) {
	createdBet, err := db.repository.Add(bet)
	if err != nil {
		return domain.BetModel{}, err
	}
	return createdBet, nil
}
