package betusecases

import (
	"log"
	"lucky-sena/app/protocols"
	"lucky-sena/domain/bet"
)

type DbAddBet struct {
	Repository protocols.AddBetRepository
}

func NewAddBet(repository protocols.AddBetRepository) *DbAddBet {
	return &DbAddBet{Repository: repository}
}

func (db *DbAddBet) AddBet(bet bet.Bet) bet.BetModel {
	createdBet, err := db.Repository.Add(bet)
	if err != nil {
		log.Fatalln(err)
	}
	return createdBet
}
