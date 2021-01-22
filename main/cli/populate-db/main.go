package main

import (
	"log"
	"lucky-sena/infra/parser"
	"lucky-sena/main/factories"
	"os"
)

func main() {
	p := factories.NewXLSXParser()
	f, ok := os.LookupEnv("XLSX_FILE")
	if !ok {
		log.Println("XLSX_FILE must not be empty")
	}

	bets := p.Parse(parser.Options{FileName: f})
	chunkSize := 100

	betsProcessed, err := persistIntoDatabaseConcurrently(bets, chunkSize)
	if err != nil {
		log.Fatalln("Something was wrong when trying to persist xlsx converted bets into db: ", err)
	}
	log.Printf("%v Bets from xlsx persisted into db successfully\n", betsProcessed)
}
