package bet

import "time"

type BetModel struct {
	ID      string    `bson:"id" json:"id"`
	Numbers []int     `bson:"numbers" json:"numbers"`
	Code    int       `bson:"code" json:"code"`
	Date    time.Time `bson:"date" json:"date"`
}

type AddBet interface {
	add(Bet) BetModel
}
