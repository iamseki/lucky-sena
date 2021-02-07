package mongodb

import (
	"lucky-sena/domain/bet"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{"code", -1}})

	cursor, err := resultsCollection.Find(a.Client.Ctx, bson.D{}, findOptions)
	if err != nil {
		defer cursor.Close(a.Client.Ctx)
		return nil, err
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

func (a *FindBetMongoRepository) FindBetsByNumbers(numbers []int) ([]bet.Bet, error) {
	return nil, nil
}
