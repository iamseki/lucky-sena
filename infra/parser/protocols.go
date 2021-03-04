package parser

import (
	"lucky-sena/domain"
)

// Parser expose the methods to be implemented by who knows how to do it
type Parser interface {
	Parse(Options) []domain.Bet
}

// Options to used to parse csv files
type Options struct {
	FileName string
}
