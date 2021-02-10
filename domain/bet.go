package domain

import (
	"time"
)

// Bet represents some bets made by someone in a specific game
type Bet struct {
	Numbers []int     `bson:"numbers" json:"numbers"`
	Code    int       `bson:"code" json:"code"`
	Date    time.Time `bson:"date" json:"date"`
}

type BetModel struct {
	ID      string    `bson:"id" json:"id"`
	Numbers []int     `bson:"numbers" json:"numbers"`
	Code    int       `bson:"code" json:"code"`
	Date    time.Time `bson:"date" json:"date"`
}

const BetCoast = 4.5

type BetCsvConverter interface {
	ConvertBetsToCsv() error
}

type BetAnalyzer interface {
	NextBetCode() int
	IsBetAlreadyWon([]int) bool
}

type BetSaver interface {
	AddBet(Bet) BetModel
}

type BetInserter interface {
	InsertBets([]Bet) error
}

type BetFinder interface {
	Find() ([]Bet, error)
}
