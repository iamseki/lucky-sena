package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo struct
type Mongo struct {
	Client *mongo.Client
	Ctx    context.Context
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
	return &Mongo{client, ctx}
}
