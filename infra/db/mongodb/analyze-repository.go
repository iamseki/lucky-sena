package mongodb

// AnalyzeBetMongoRepository represents a struct who knows how to analyze bets in database
type AnalyzeBetMongoRepository struct {
	FindBetMongoRepository
}

// NewAnalyzeBetMongoRepository returns an instance of AnalyzeBetMongoRepository struct
func NewAnalyzeBetMongoRepository(collection string) *AnalyzeBetMongoRepository {
	return &AnalyzeBetMongoRepository{
		FindBetMongoRepository{Client: newMongoConnection(), Collection: collection},
	}
}
