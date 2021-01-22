package main

import (
	"log"
	"lucky-sena/domain/bet"
	"lucky-sena/main/factories"
	"sync"
)

func persistIntoDatabaseConcurrently(betsToPersist []bet.Bet, chunkSize int) (int, error) {
	var wg sync.WaitGroup
	var betsProcessed int

	insertUseCase := factories.NewInsertBetsUseCase()

	for {
		betsToProcess := betsToPersist[:chunkSize]
		betsToPersist = betsToPersist[chunkSize:]

		betsProcessed += len(betsToProcess)
		wg.Add(1)
		go func(w *sync.WaitGroup, bets []bet.Bet) {
			defer w.Done()
			insertUseCase.InsertBets(bets)
		}(&wg, betsToProcess)

		remainingBets := len(betsToPersist)
		if remainingBets == 0 {
			break
		}
		if remainingBets < chunkSize {
			chunkSize = remainingBets
		}
	}
	wg.Wait()

	return betsProcessed, nil
}

func persistIntoDatabase(betsToPersist []bet.Bet, chunkSize int) (int, error) {
	betsProcessed := len(betsToPersist)
	log.Println("Instantiate insert use case")
	insertUseCase := factories.NewInsertBetsUseCase()
	log.Println("Executing insert use case")
	err := insertUseCase.InsertBets(betsToPersist)
	if err != nil {
		return 0, err
	}

	return betsProcessed, nil
}
