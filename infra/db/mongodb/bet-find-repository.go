package mongodb

import (
	"lucky-sena/domain/bet"

	"go.mongodb.org/mongo-driver/bson"
)

type FindBetMongoRepository struct {
	Client *Mongo
}

func NewFindBetMongoRepository() *FindBetMongoRepository {
	return &FindBetMongoRepository{
		Client: newMongoConnection(),
	}
}

func (a *FindBetMongoRepository) Find() ([]bet.Bet, error) {
	resultsCollection := a.Client.getCollection("results")
	cursor, err := resultsCollection.Find(a.Client.Ctx, bson.D{})
	if err != nil {
		return nil, err
		defer cursor.Close(a.Client.Ctx)
	}

	var bets []bet.Bet
	for cursor.Next(a.Client.Ctx) {
		var bet bet.Bet
		cursor.Decode(&bet)
		bets = append(bets, bet)
	}
	return bets, nil
}

func (a *FindBetMongoRepository) FindBetByCode(code int) (bet.Bet, error) {
	resultsCollection := a.Client.getCollection("results")
	var bet bet.Bet
	resultsCollection.FindOne(a.Client.Ctx, bson.M{"code": code}).Decode(bet)
	return bet, nil
}
