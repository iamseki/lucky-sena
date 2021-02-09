package parser

import (
	"lucky-sena/domain"
)

type Parser interface {
	Parse(Options) []domain.Bet
}

type Options struct {
	FileName string
}
