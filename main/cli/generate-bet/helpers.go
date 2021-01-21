package main

import (
	"flag"
	"log"
	"lucky-sena/infra/generator"
)

func parseFlags(f *flags) {
	flag.IntVar(&f.bets, "b", 0, "b is equal the number of bets to be done")
	flag.StringVar(&f.excludedNumbersCSV, "e", "", "e is the numbers to exclude in csv format: 1,2,3,4,5,6")
	flag.IntVar(&f.gameCode, "c", 0, "number of the game to make a bet")
	flag.BoolVar(&f.persist, "p", false, "persist to database if true")

	flag.Parse()

	requiredFlags := []string{"b", "e"}

	if !isRequiredFlagsPassed(requiredFlags) {
		log.Fatalln(`
******************* THIS PROGRAM IS DESIGNED TO RUN WITH THE FOLLOWING FLAG's OPTIONS **********
------->   ./generate-bet --e=1,2,3 --b=2 --c=2334`)
	}
}

func isRequiredFlagsPassed(requiredFlags []string) bool {
	for _, name := range requiredFlags {
		if !isFlagPassed(name) {
			return false
		}
	}
	return true
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func isFlagOptional(name string) bool {
	return name == "c" || name == "p"
}

func generateBets(f *flags) []generator.GenaretedBet {
	log.Printf("Generating bets to game code: %v\n", f.gameCode)
	excludedBalls := generator.ConvertCSVIntoIntSlice(f.excludedNumbersCSV)
	gen := generator.Factory(generator.Default)
	return gen.Generate(generator.Options{BetsToGenerate: f.bets, ExcludedNumbers: excludedBalls})
}
