package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain/bet"
)

type DbInsertBets struct {
	Repository protocols.InsertBetsRepository
}

func NewInsertBets(repository protocols.InsertBetsRepository) *DbInsertBets {
	return &DbInsertBets{Repository: repository}
}

func (db *DbInsertBets) InsertBets(bets []bet.Bet) error {
	err := db.Repository.InsertMany(bets)
	if err != nil {
		return err
	}
	return nil
}
