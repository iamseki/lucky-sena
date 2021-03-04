package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain"
)

type DbInsertBets struct {
	repository protocols.InsertBetsRepository
}

func NewInsertBets(repository protocols.InsertBetsRepository) *DbInsertBets {
	return &DbInsertBets{repository}
}

func (db *DbInsertBets) InsertBets(bets []domain.Bet) error {
	err := db.repository.InsertMany(bets)
	if err != nil {
		return err
	}
	return nil
}
