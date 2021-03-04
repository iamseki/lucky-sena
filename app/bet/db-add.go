package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain"
)

// DbAddBet represents a struct that implements domain.BetSaver interface
type DbAddBet struct {
	repository protocols.AddBetRepository
}

// NewAddBet returns an instance of DbAddBet struct
func NewAddBet(repository protocols.AddBetRepository) *DbAddBet {
	return &DbAddBet{repository}
}

// AddBet adds a Bet into database
func (db *DbAddBet) AddBet(bet domain.Bet) (domain.BetModel, error) {
	createdBet, err := db.repository.Add(bet)
	if err != nil {
		return domain.BetModel{}, err
	}
	return createdBet, nil
}
