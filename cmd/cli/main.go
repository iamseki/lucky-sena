package main

import (
	"flag"
	"log"
	"lucky-sena/app/bet"
	"lucky-sena/domain/bet"
	"lucky-sena/infra/db/mongodb"
	"lucky-sena/infra/generator"
	"lucky-sena/infra/parser"
	"sync"
	"time"
)

func main() {
	var bets int
	var gameCode int
	var excludedNumbersCSV string
	var filename string
	var wg sync.WaitGroup

	flag.IntVar(&bets, "b", 0, "b is equal the number of bets to be done")
	flag.StringVar(&excludedNumbersCSV, "e", "", "e is the numbers to exclude in csv format: 1,2,3,4,5,6")
	flag.StringVar(&filename, "f", "", "the name of xlsx file to read and inserts into database")
	flag.IntVar(&gameCode, "c", 0, "number of the game to make a bet")

	flag.Parse()

	if bets > 0 {
		log.Printf("Generating bets to game code: %v\n", gameCode)
		excludedBalls := generator.ConvertCSVIntoIntSlice(excludedNumbersCSV)
		gen := generator.Factory(generator.Default)
		betsGenerated := gen.Generate(generator.Options{BetsToGenerate: bets, ExcludedNumbers: excludedBalls})
		log.Println(betsGenerated)

		log.Println("Persisting generated bets into database")
		addBetRepository := mongodb.NewAddBetMongoRepository()
		addBet := betusecases.NewAddBet(addBetRepository)

		wg.Add(len(betsGenerated))
		for _, b := range betsGenerated {
			go func(b generator.GenaretedBet) {
				defer wg.Done()
				addBet.Add(bet.Bet{Numbers: b.Numbers, Code: gameCode, Date: time.Now()})
			}(b)
		}
		wg.Wait()
	} else if filename != "" {
		/*
			xw := services.Factory(services.WriteXlsxResults).(*services.XlsxWriter)
			xw.SetFileName(filename)
			xw.Run()
		*/
		p := parser.Factory(parser.Xlsx)
		bets := p.Parse(parser.Options{filename})
		log.Println(bets)
	} else {
		log.Fatalln(`
******************* THIS PROGRAM IS DESIGNED TO RUN WITH THE FOLLOWING FLAG's OPTIONS **********
./lucky-sena --e=1,2,3 --b=2
./lucky-sena --f="mega_sena_asloterias_ate_concurso_2325_sorteio.xlsx"
		`)
	}
}
