package main

import (
	"lucky-sena/infra/generator"
)

type generateBetsHandler func(bets int, excludeds []int) []generator.GenaretedBet

func generateBetHandler(bets int, excludeds []int) []generator.GenaretedBet {
	gen := generator.Factory(generator.Default)
	return gen.Generate(generator.Options{BetsToGenerate: bets, ExcludedNumbers: excludeds})
}
