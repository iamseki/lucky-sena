package main

import (
	"log"
	"lucky-sena/domain"
	"lucky-sena/infra/generator"
	"lucky-sena/main/factories"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	options := &flags{}
	parseFlags(options)
	if options.gameCode == 0 {
		analyzeBet := factories.NewAnalyzeBetUseCase()
		options.gameCode = analyzeBet.NextBetCode()
	}

	addBetUseCase := factories.NewAddBetUseCase()
	betsGenerated := generateBets(options)

	if options.persist {
		wg.Add(len(betsGenerated))
		log.Println("Persisting generated bets into database")
		for _, b := range betsGenerated {
			go func(b generator.GenaretedBet) {
				defer wg.Done()
				addBetUseCase.AddBet(domain.Bet{Numbers: b.Numbers, Code: options.gameCode, Date: time.Now()})
			}(b)
		}
		wg.Wait()
		log.Println("Bets persisted into database !")
	} else {
		log.Println("Bets generated: \n", betsGenerated)
	}
}
