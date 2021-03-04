package mongodb

import (
	"lucky-sena/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindBetMongoRepository knows how to find bets in MongoDB database
type FindBetMongoRepository struct {
	Client     *Mongo
	Collection string
}

// NewFindBetMongoRepository returns an instance of FindBetMongoRepository
func NewFindBetMongoRepository(collection string) *FindBetMongoRepository {
	return &FindBetMongoRepository{
		Client:     newMongoConnection(),
		Collection: collection,
	}
}

// Find all bets in database
func (repo *FindBetMongoRepository) Find() ([]domain.Bet, error) {
	resultsCollection := repo.Client.getCollection(repo.Collection)

	findOptions := options.Find()
	// Sort by `price` field descending
	findOptions.SetSort(bson.D{{"code", -1}})

	cursor, err := resultsCollection.Find(repo.Client.Ctx, bson.D{}, findOptions)
	if err != nil {
		//defer cursor.Close(repo.Client.Ctx)
		return nil, err
	}

	var bets []domain.Bet
	for cursor.Next(repo.Client.Ctx) {
		var bet domain.Bet
		cursor.Decode(&bet)
		bets = append(bets, bet)
	}
	return bets, nil
}

// FindBetByCode find bets by game code in results collection
func (repo *FindBetMongoRepository) FindBetByCode(code int) (domain.Bet, error) {
	resultsCollection := repo.Client.getCollection("results")
	var bet domain.Bet
	resultsCollection.FindOne(repo.Client.Ctx, bson.M{"code": code}).Decode(bet)
	return bet, nil
}

// FindBetsByNumbers find bets by numbers passed through params
func (repo *FindBetMongoRepository) FindBetsByNumbers(numbers []int) ([]domain.Bet, error) {
	return nil, nil
}
