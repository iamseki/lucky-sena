package main

import (
	"errors"
	"flag"
	"log"
	"lucky-sena/domain"
	"lucky-sena/main/factories"
	"regexp"
	"sync"
)

func validateFileExtension(filename string) error {
	validExtension := regexp.MustCompile(`^.*\.(csv|xlsx)$`)

	if isValid := validExtension.MatchString(filename); !isValid {
		return errors.New("invalid filename extension -> recieved: " + filename + " expected: *.(csv|xlsx)")
	}
	return nil
}

func parseFlags(f *flags) {
	flag.BoolVar(&f.concurrency, "concurrency", false, "persist to database if true")
	flag.IntVar(&f.chunkSize, "chunk", 100, "chunks to be processed")
	flag.StringVar(&f.collection, "collection", "", "collection to be read")
	flag.StringVar(&f.extension, "ext", "csv", "extension file to get the bets")

	flag.Parse()
}

func persist(fn persistFunction, bets []domain.Bet, options *flags) {
	betsProcessed, err := fn(bets, options.collection, options.chunkSize)
	if err != nil {
		log.Fatalln("Something was wrong when trying to persist xlsx converted bets into db: ", err)
	}
	log.Printf("%v Bets from xlsx persisted into db successfully\n", betsProcessed)
}

func persistIntoDatabaseConcurrently(betsToPersist []domain.Bet, collection string, chunkSize int) (int, error) {
	var wg sync.WaitGroup
	var betsProcessed int

	insertUseCase := factories.NewInsertBetsUseCase(collection)

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

func persistIntoDatabase(betsToPersist []domain.Bet, collection string, chunkSize int) (int, error) {
	betsProcessed := len(betsToPersist)
	log.Println("Instantiate insert use case")
	insertUseCase := factories.NewInsertBetsUseCase(collection)
	log.Println("Executing insert use case")
	err := insertUseCase.InsertBets(betsToPersist)
	if err != nil {
		return 0, err
	}

	return betsProcessed, nil
}
