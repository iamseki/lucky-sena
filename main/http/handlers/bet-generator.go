package handlers

import (
	"lucky-sena/infra/generator"
)

// RequestBetsHandle represents the body of request expected in that handler
type RequestBetsHandle struct {
	Bets         int   `json:"bets"`
	ExcludedBets []int `json:"excludeds"`
}

// ResponseBetsHandle represents the body of response returned from that handler
type ResponseBetsHandle struct {
	Bets []generator.GenaretedBet `json:"bets"`
}

// GenerateBetsHandler is the func expected to recieve in adapter
type GenerateBetsHandler func(bets int, excludeds []int) []generator.GenaretedBet

// GenerateBetsHandle is the implementation of GenerateBetsHandler func
func GenerateBetsHandle(bets int, excludeds []int) []generator.GenaretedBet {
	gen := generator.Factory(generator.Default)
	return gen.Generate(generator.Options{BetsToGenerate: bets, ExcludedNumbers: excludeds})
}
