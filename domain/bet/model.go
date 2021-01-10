package bet

import (
	"time"
)

// Bet represents some bets made by someone in a specific game
type Bet struct {
	Numbers []int     `bson:"numbers json:"numbers""`
	Code    int       `bson:"code" json:"code"`
	Date    time.Time `bson:"date" json:"date"`
	Coast   float32   `bson:"coast" json:"coast"`
}

const Coast = 4.5
