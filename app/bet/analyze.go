package betusecases

import (
	"lucky-sena/app/protocols"
)

// AnalyzeBet represents a struct that implements domain.BetAnalyzer interface
type AnalyzeBet struct {
	betRepository protocols.AnalyzeBetRepository
}

// NewAnalyzeBet returns an instance of AnalyzeBet struct
func NewAnalyzeBet(repository protocols.AnalyzeBetRepository) *AnalyzeBet {
	return &AnalyzeBet{repository}
}

// NextBetCode returns the next code of the latest bet registered in database
func (analyze *AnalyzeBet) NextBetCode() (int, error) {
	bets, err := analyze.betRepository.Find()
	if err != nil {
		return -1, err
	}
	return len(bets) + 1, nil
}

// IsBetAlreadyWon checks if the numbers inputed has won the MegaSena somehow
func (analyze *AnalyzeBet) IsBetAlreadyWon(numbers []int) (bool, error) {
	bets, err := analyze.betRepository.FindBetsByNumbers(numbers)
	if err != nil {
		return false, err
	}
	return len(bets) > 0, nil
}
