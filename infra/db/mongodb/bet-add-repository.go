package mongodb

import (
	"fmt"
	"lucky-sena/domain"
)

// AddBetMongoRepository knows how to insert bets into MongoDB database
type AddBetMongoRepository struct {
	Client *Mongo
}

// NewAddBetMongoRepository returns an instance of AddBetMongoRepository
func NewAddBetMongoRepository() *AddBetMongoRepository {
	return &AddBetMongoRepository{
		Client: newMongoConnection(),
	}
}

// Add just add the bet into mongodb
func (m *AddBetMongoRepository) Add(betToAdd domain.Bet) (domain.BetModel, error) {
	collectionSena := m.Client.getCollection("bets")
	result, err := collectionSena.InsertOne(m.Client.Ctx, betToAdd)
	if err != nil {
		return domain.BetModel{}, err
	}
	id := fmt.Sprintf("%v", result.InsertedID)

	return domain.BetModel{ID: id, Numbers: betToAdd.Numbers, Code: betToAdd.Code, Date: betToAdd.Date}, nil
}
