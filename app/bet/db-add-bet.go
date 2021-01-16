package betusecases

import (
	"log"
	"lucky-sena/domain/bet"
)

type DbAddBet struct {
	Repository AddBetRepository
}

func NewAddBet(repository AddBetRepository) *DbAddBet {
	return &DbAddBet{Repository: repository}
}

func (db *DbAddBet) AddBet(bet bet.Bet) bet.BetModel {
	createdBet, err := db.Repository.Add(bet)
	if err != nil {
		log.Fatalln(err)
	}
	return createdBet
}
