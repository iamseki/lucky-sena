package handlers

import (
	"lucky-sena/infra/generator"
)

type GenerateBetsHandler func(bets int, excludeds []int) []generator.GenaretedBet

func GenerateBetsHandle(bets int, excludeds []int) []generator.GenaretedBet {
	gen := generator.Factory(generator.Default)
	return gen.Generate(generator.Options{BetsToGenerate: bets, ExcludedNumbers: excludeds})
}
