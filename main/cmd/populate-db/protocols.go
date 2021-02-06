package main

import "lucky-sena/domain/bet"

type persistFunction func(betsToPersist []bet.Bet, chunkSize int) (int, error)

type flags struct {
	concurrency bool
	chunkSize   int
}
