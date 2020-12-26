package main

import (
	"flag"
	"log"
	"lucky-sena/services"
)

func main() {
	// Flags variables
	var bets int
	var excludedNumbers string
	var filename string

	flag.IntVar(&bets, "b", 0, "b is equal the number of bets to be done")
	flag.StringVar(&excludedNumbers, "e", "", "e is the numbers to exclude in csv format: 1,2,3,4,5,6")
	flag.StringVar(&filename, "f", "", "the name of xlsx file to read and inserts into database")

	// Insert the flags into those variables
	flag.Parse()

	if bets > 0 { // Do the bets generator stuff
		// Type assertion is ok because BetGenerator Implements IFactory
		bg := services.Factory(services.BetsGen).(*services.BetGenerator)
		bg.SetBet(bets)
		bg.SetExcludedNumbers(excludedNumbers)
		// Generate the random numbers and prints it
		bg.Run()
	} else if filename != "" {
		xw := services.Factory(services.WriteXlsxResults).(*services.XlsxWriter)
		xw.SetFileName(filename)
		// Read xlsx file
		xw.Run()
	} else {
		log.Fatalln(`
******************* THIS PROGRAM IS DESIGNED TO RUN WITH THE FOLLOWING FLAG's OPTIONS **********
./lucky-sena --e=1,2,3 --b=2
./lucky-sena --f="mega_sena_asloterias_ate_concurso_2325_sorteio.xlsx"
		`)
	}
}
