package mongodb

import (
	"fmt"
	"lucky-sena/domain/bet"
)

type AddBetMongoRepository struct {
	Client *Mongo
}

func NewAddBetMongoRepository() *AddBetMongoRepository {
	return &AddBetMongoRepository{
		Client: newMongoConnection(),
	}
}

func (m *AddBetMongoRepository) Add(betToAdd bet.Bet) (bet.BetModel, error) {
	collectionSena := m.Client.getCollection("bets")
	result, err := collectionSena.InsertOne(m.Client.Ctx, betToAdd)
	if err != nil {
		return bet.BetModel{}, err
	}
	id := fmt.Sprintf("%v", result.InsertedID)

	return bet.BetModel{ID: id, Numbers: betToAdd.Numbers, Code: betToAdd.Code, Date: betToAdd.Date}, nil
}
