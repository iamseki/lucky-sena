package main

import (
	"flag"
	"log"
	"lucky-sena/main/factories"
)

func main() {
	var collection string
	flag.StringVar(&collection, "collection", "results", "collection to be read")
	flag.StringVar(&collection, "c", "results", "shortly flag to what collection to be read")

	if collection == "results" {
		c := factories.NewResultsBetToCsvConverter()
		c.ConvertBetsToCsv()
		log.Println("Results collection converted into csv file successfully")
	} else if collection == "bets" {
		c := factories.NewMadeBetToCsvConverter()
		c.ConvertBetsToCsv()
		log.Println("Bets collection converted into csv file successfully")
	} else {
		log.Printf("Collection %v doesnt exist\n", collection)
	}

}
