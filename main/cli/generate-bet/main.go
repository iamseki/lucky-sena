package main

import (
	"log"
	"lucky-sena/domain/bet"
	"lucky-sena/infra/generator"
	"lucky-sena/main/factories"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	f := &flags{}
	parseFlags(f)
	if f.gameCode == 0 {
		analyzeBet := factories.NewAnalyzeBetUseCase()
		f.gameCode = analyzeBet.NextBetCode()
	}

	addBetUseCase := factories.NewAddBetUseCase()
	betsGenerated := generateBets(f)

	if f.persist {
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
	} else {
		log.Println("Bets generated: \n", betsGenerated)
	}
}
