package mongodb

type AnalyzeBetMongoRepository struct {
	FindBetMongoRepository
}

func NewAnalyzeBetMongoRepository(collection string) *AnalyzeBetMongoRepository {
	return &AnalyzeBetMongoRepository{
		FindBetMongoRepository{Client: newMongoConnection(), Collection: collection},
	}
}
