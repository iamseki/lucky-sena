package bet

import (
	"time"
)

// Bet represents some bets made by someone in a specific game
type Bet struct {
	Numbers []Balls   `bson:"numbers json:"numbers""`
	Code    int       `bson:"code" json:"code"`
	Date    time.Time `bson:"date" json:"date"`
	Coast   float32   `bson:"coast" json:"coast"`
}

// Balls represents an array with the numbers of a bet
type Balls struct {
	Bet []int `bson:"bet" json:"bet"`
}
