package betusecases

import (
	"log"
)

type AnalyzeBet struct {
	betRepository AnalyzeBetRepository
}

func NewAnalyzeBet(repository AnalyzeBetRepository) *AnalyzeBet {
	return &AnalyzeBet{repository}
}

func (analyze *AnalyzeBet) NextBetCode() int {
	bets, err := analyze.betRepository.Find()
	if err != nil {
		log.Fatalln(err)
	}
	return len(bets) + 1
}

func (analyze *AnalyzeBet) IsBetAlreadyWon(numbers []int) bool {
	bets, err := analyze.betRepository.FindBetsByNumbers(numbers)
	if err != nil {
		log.Fatalln(err)
	}
	return len(bets) > 0
}
