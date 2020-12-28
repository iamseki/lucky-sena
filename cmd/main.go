package main

import (
	"flag"
	"log"
	"lucky-sena/services"
)

func main() {
	var bets int
	var excludedNumbers string
	var filename string

	flag.IntVar(&bets, "b", 0, "b is equal the number of bets to be done")
	flag.StringVar(&excludedNumbers, "e", "", "e is the numbers to exclude in csv format: 1,2,3,4,5,6")
	flag.StringVar(&filename, "f", "", "the name of xlsx file to read and inserts into database")

	flag.Parse()

	if bets > 0 {
		betGenerator := services.Factory(services.BetsGen).(*services.BetGenerator)
		betGenerator.SetNumberOfBets(bets)
		betGenerator.SetExcludedNumbers(excludedNumbers)
		betGenerator.Run()
	} else if filename != "" {
		xw := services.Factory(services.WriteXlsxResults).(*services.XlsxWriter)
		xw.SetFileName(filename)
		xw.Run()
	} else {
		log.Fatalln(`
******************* THIS PROGRAM IS DESIGNED TO RUN WITH THE FOLLOWING FLAG's OPTIONS **********
./lucky-sena --e=1,2,3 --b=2
./lucky-sena --f="mega_sena_asloterias_ate_concurso_2325_sorteio.xlsx"
		`)
	}
}
