package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Result represents a single result of a game document
type Result struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Bet  []int              `bson:"bet"`
	Code int                `bson:"code"`
	Date time.Time          `bson:"date"`
}
