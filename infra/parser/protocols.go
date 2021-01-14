package parser

import "lucky-sena/domain/bet"

type Parser interface {
	Parse(Options) []bet.Bet
}

type Options struct {
	FileName string
}
