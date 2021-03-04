package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain"
)

// DbInsertBets represents a struct that implements domain.BetInserter interface
type DbInsertBets struct {
	repository protocols.InsertBetsRepository
}

// NewInsertBets returns an instance of DbInsertBets struct
func NewInsertBets(repository protocols.InsertBetsRepository) *DbInsertBets {
	return &DbInsertBets{repository}
}

// InsertBets insert whole bets passed through params into database
func (db *DbInsertBets) InsertBets(bets []domain.Bet) error {
	err := db.repository.InsertMany(bets)
	if err != nil {
		return err
	}
	return nil
}
