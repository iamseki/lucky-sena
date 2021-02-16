package main

import (
	"log"
	"lucky-sena/infra/parser"
	"lucky-sena/main/factories"
	"os"
)

func main() {
	p := factories.NewXLSXParser()
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
	bets := p.Parse(parser.Options{FileName: f})

	if options.concurrency {
		persist(persistIntoDatabaseConcurrently, bets, options.chunkSize)
	} else {
		persist(persistIntoDatabase, bets, options.chunkSize)
	}
}
