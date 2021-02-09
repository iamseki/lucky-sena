package main

import "lucky-sena/domain"

type persistFunction func(betsToPersist []domain.Bet, chunkSize int) (int, error)

type flags struct {
	concurrency bool
	chunkSize   int
}
