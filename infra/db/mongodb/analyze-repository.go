package mongodb

type AnalyzeBetMongoRepository struct {
	FindBetMongoRepository
}

func NewAnalyzeBetMongoRepository() *AnalyzeBetMongoRepository {
	return &AnalyzeBetMongoRepository{
		FindBetMongoRepository{Client: newMongoConnection()},
	}
}
