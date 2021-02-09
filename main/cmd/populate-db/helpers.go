package main

import (
	"flag"
	"log"
	"lucky-sena/domain"
	"lucky-sena/main/factories"
	"sync"
)

func parseFlags(f *flags) {
	flag.BoolVar(&f.concurrency, "concurrency", false, "persist to database if true")
	flag.IntVar(&f.chunkSize, "chunk", 100, "chunks to be processed")
}

func persist(fn persistFunction, bets []domain.Bet, chunk int) {
	betsProcessed, err := fn(bets, chunk)
	if err != nil {
		log.Fatalln("Something was wrong when trying to persist xlsx converted bets into db: ", err)
	}
	log.Printf("%v Bets from xlsx persisted into db successfully\n", betsProcessed)
}

func persistIntoDatabaseConcurrently(betsToPersist []domain.Bet, chunkSize int) (int, error) {
	var wg sync.WaitGroup
	var betsProcessed int

	insertUseCase := factories.NewInsertBetsUseCase()

	for {
		betsToProcess := betsToPersist[:chunkSize]
		betsToPersist = betsToPersist[chunkSize:]

		betsProcessed += len(betsToProcess)
		wg.Add(1)
		go func(w *sync.WaitGroup, bets []domain.Bet) {
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

func persistIntoDatabase(betsToPersist []domain.Bet, chunkSize int) (int, error) {
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
