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

// BetModel represents the bet returned from database
type BetModel struct {
	ID      string    `bson:"id" json:"id"`
	Numbers []int     `bson:"numbers" json:"numbers"`
	Code    int       `bson:"code" json:"code"`
	Date    time.Time `bson:"date" json:"date"`
}

// BetCoast is the price of a single bet made in MegaSena
const BetCoast = 4.5

// BetCsvConverter expose the methods to be implemented by some code that knows how to do it
type BetCsvConverter interface {
	ConvertBetsToCsv() error
}

// BetAnalyzer expose the methods to be implemented by some code that knows how to do it
type BetAnalyzer interface {
	NextBetCode() (int, error)
	IsBetAlreadyWon([]int) (bool, error)
}

// BetSaver expose the methods to be implemented by some code that knows how to do it
type BetSaver interface {
	AddBet(Bet) (BetModel, error)
}

// BetInserter expose the methods to be implemented by some code that knows how to do it
type BetInserter interface {
	InsertBets([]Bet) error
}

// BetFinder expose the methods to be implemented by some code that knows how to do it
type BetFinder interface {
	Find() ([]Bet, error)
}
