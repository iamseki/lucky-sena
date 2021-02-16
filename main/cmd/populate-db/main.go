package main

import (
	"log"
	"lucky-sena/domain"
	"lucky-sena/infra/parser"
	"lucky-sena/main/factories"
	"os"
)

func main() {
	f, ok := os.LookupEnv("PARSE_FILE")
	if !ok {
		log.Fatalln("PARSE_FILE must not be empty")
	}

	err := validateFileExtension(f)
	if err != nil {
		log.Fatalln(err)
	}

	options := &flags{}
	parseFlags(options)

	var bets []domain.Bet
	if options.extension == "csv" {
		p := factories.NewCSVParser()
		bets = p.Parse(parser.Options{FileName: f})
	} else if options.extension == "xlsx" {
		p := factories.NewXLSXParser()
		bets = p.Parse(parser.Options{FileName: f})
	}

	if options.concurrency {
		persist(persistIntoDatabaseConcurrently, bets, options)
	} else {
		persist(persistIntoDatabase, bets, options)
	}
}
