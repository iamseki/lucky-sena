package main

import (
	"log"
	betusecases "lucky-sena/app/bet"
	"lucky-sena/infra/db/mongodb"
)

func main() {
	analyzeBetRepository := mongodb.NewAnalyzeBetMongoRepository()
	analyzeBet := betusecases.NewAnalyzeBet(analyzeBetRepository)

	nextCode := analyzeBet.NextBetCode()
	log.Println(nextCode)

	findBetRepository := mongodb.NewFindBetMongoRepository()
	findBet := betusecases.NewFindBet(findBetRepository)

	bets, _ := findBet.Find()
	log.Println(bets)
}
