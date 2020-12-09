package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Result represents a single result document
type Result struct {
	ID   primitive.ObjectID `bson:"_id"`
	Bet  []int              `bson:"bet"`
	Code int                `bson:"code"`
	Date time.Time          `bson:"created_at"`
}
