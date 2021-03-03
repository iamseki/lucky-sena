package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Client   *mongo.Client
	Ctx      context.Context
	database string
}

var mongoClient *mongo.Client

func newMongoConnection() *Mongo {
	mongoConn, declared := os.LookupEnv("MONGO_URI")
	if !declared {
		mongoConn = "mongodb://localhost:27017"
	}

	ctx := getMongoConnection(mongoConn)
	return &Mongo{mongoClient, ctx, "sena"}
}

func (m *Mongo) getCollection(name string) *mongo.Collection {
	return m.Client.Database(m.database).Collection(name)
}

func getMongoConnection(uri string) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if mongoClient == nil {
		mongoClient, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri), options.Client())
		err := mongoClient.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Println("Retrying to connect to mongodb")
			mongoClient = nil
			getMongoConnection(uri)
		} else {
			log.Println("Connected into mongodb: ", uri)
		}
		return ctx
	}
	log.Println("Using active mongo connection")
	return ctx
}
