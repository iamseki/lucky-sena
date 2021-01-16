package main

import (
	"log"
	"lucky-sena/app/bet"
	"lucky-sena/domain/bet"
	"lucky-sena/infra/db/mongodb"
	"lucky-sena/infra/generator"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	f := &flags{}
	parseFlags(f)

	addBetRepository := mongodb.NewAddBetMongoRepository()
	addBetUseCase := betusecases.NewAddBet(addBetRepository)

	betsGenerated := makeBets(f)
	wg.Add(len(betsGenerated))

	log.Println("Persisting generated bets into database")
	for _, b := range betsGenerated {
		go func(b generator.GenaretedBet) {
			defer wg.Done()
			addBetUseCase.AddBet(bet.Bet{Numbers: b.Numbers, Code: f.gameCode, Date: time.Now()})
		}(b)
	}
	wg.Wait()
	log.Println("Bets persisted into database !")
}
