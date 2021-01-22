package mongodb

import (
	"lucky-sena/domain/bet"
)

type InsertBetsMongoRepository struct {
	Client *Mongo
}

func NewInsertBetsMongoRepository() *InsertBetsMongoRepository {
	return &InsertBetsMongoRepository{
		Client: newMongoConnection(),
	}
}

func (m *InsertBetsMongoRepository) InsertMany(bets []bet.Bet) error {
	resultCollection := m.Client.getCollection("results")

	docs := []interface{}{}
	for _, bet := range bets {
		docs = append(docs, bet)
	}

	_, err := resultCollection.InsertMany(m.Client.Ctx, docs)
	if err != nil {
		return err
	}
	return nil
}
