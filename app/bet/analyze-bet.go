package betusecases

import (
	"fmt"
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

func (analyze *AnalyzeBet) IsBetAlreadyWon(code int) bool {
	bet, err := analyze.betRepository.FindBetByCode(code)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bet)
	return true
}
