package mongodb

import (
	"fmt"
	"lucky-sena/domain"
)

type AddBetMongoRepository struct {
	Client *Mongo
}

func NewAddBetMongoRepository() *AddBetMongoRepository {
	return &AddBetMongoRepository{
		Client: newMongoConnection(),
	}
}

func (m *AddBetMongoRepository) Add(betToAdd domain.Bet) (domain.BetModel, error) {
	collectionSena := m.Client.getCollection("bets")
	result, err := collectionSena.InsertOne(m.Client.Ctx, betToAdd)
	if err != nil {
		return domain.BetModel{}, err
	}
	id := fmt.Sprintf("%v", result.InsertedID)

	return domain.BetModel{ID: id, Numbers: betToAdd.Numbers, Code: betToAdd.Code, Date: betToAdd.Date}, nil
}
