package mongodb

import (
	"lucky-sena/domain"
)

// InsertBetsMongoRepository knows how to insert bets into mongodb
type InsertBetsMongoRepository struct {
	Client     *Mongo
	Collection string
}

// NewInsertBetsMongoRepository returns an instance of InsertBetsMongoRepository
func NewInsertBetsMongoRepository(collection string) *InsertBetsMongoRepository {
	return &InsertBetsMongoRepository{
		Client:     newMongoConnection(),
		Collection: collection,
	}
}

// InsertMany insert whole bets into mongo collection
func (m *InsertBetsMongoRepository) InsertMany(bets []domain.Bet) error {
	resultCollection := m.Client.getCollection(m.Collection)

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
