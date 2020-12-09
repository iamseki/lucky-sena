package main

import (
	"lucky-sena/services"
)

func main() {

	// type assertion is ok because BetGenerator Implements IFactory
	bg := services.Factory(services.BetsGen).(*services.BetGenerator)

	xw := services.Factory(services.WriteXlsxResults)

	bg.Run()
	xw.Run()
}
