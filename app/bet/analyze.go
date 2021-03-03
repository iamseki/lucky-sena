package betusecases

import (
	"lucky-sena/app/protocols"
)

type AnalyzeBet struct {
	betRepository protocols.AnalyzeBetRepository
}

func NewAnalyzeBet(repository protocols.AnalyzeBetRepository) *AnalyzeBet {
	return &AnalyzeBet{repository}
}

func (analyze *AnalyzeBet) NextBetCode() (int, error) {
	bets, err := analyze.betRepository.Find()
	if err != nil {
		return -1, err
	}
	return len(bets) + 1, nil
}

func (analyze *AnalyzeBet) IsBetAlreadyWon(numbers []int) (bool, error) {
	bets, err := analyze.betRepository.FindBetsByNumbers(numbers)
	if err != nil {
		return false, err
	}
	return len(bets) > 0, nil
}
