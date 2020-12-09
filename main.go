package main

import (
	"flag"
	"lucky-sena/services"
)

func main() {
	// GET THE FLAGS OF PROGRAM
	var bets int
	var excludedNumbers string

	flag.IntVar(&bets, "b", 0, "b is equal the number of bets to be done")
	flag.StringVar(&excludedNumbers, "e", "", "e is the numbers to exclude in csv format: 1,2,3,4,5,6")
	flag.Parse()

	// type assertion is ok because BetGenerator Implements IFactory
	bg := services.Factory(services.BetsGen).(*services.BetGenerator)

	bg.SetBet(bets)
	bg.SetExcludedNumbers(excludedNumbers)

	bg.Run()
}
