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
	// defer cancel()

	mongoConn, declared := os.LookupEnv("MONGO_URI")
	if !declared {
		log.Fatalln("MONGO_URI must be provided")
	}
	log.Printf("Trying to connect in %v\n", mongoConn)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConn))
	if err != nil {
		panic(err)
	}

	log.Printf("App is connected to MongoDB !")
	return &Mongo{client, ctx, "sena"}
}

func (m *Mongo) getCollection(name string) *mongo.Collection {
	return m.Client.Database(m.database).Collection(name)
}
