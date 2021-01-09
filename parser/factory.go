package parser

type ParserType string

const (
	Default      ParserType = "default"
	CustomParser            = "customParser"
)

func Factory(p ParserType) Parser {
	switch p {
	case Default:
		return newDefaultParser()
	case CustomParser:
		return newDefaultParser()
	default:
		return newDefaultParser()
	}
}
