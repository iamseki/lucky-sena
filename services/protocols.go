package services

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Bet represents some bets made by someone in a specific game
type Bet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Numbers []Balls            `bson:"numbers"`
	Code    int                `bson:"code"`
	Date    time.Time          `bson:"date"`
	Coast   float32            `bson:"coast"`
}

// Balls represents an array with the numbers of a bet
type Balls struct {
	Bet []int `bson:"bet"`
}

// Result represents a single result of a game document
type Result struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Bet  []int              `bson:"bet"`
	Code int                `bson:"code"`
	Date time.Time          `bson:"date"`
}
