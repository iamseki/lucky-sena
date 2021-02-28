package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client   *mongo.Client
	Ctx      context.Context
	database string
}

func newMongoConnection() *Mongo {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	mongoConn, declared := os.LookupEnv("MONGO_URI")
	if !declared {
		mongoConn = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConn))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("New client connection at:", mongoConn)
	return &Mongo{client, ctx, "sena"}
}

func (m *Mongo) getCollection(name string) *mongo.Collection {
	return m.Client.Database(m.database).Collection(name)
}
