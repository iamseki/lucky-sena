package factories

import "lucky-sena/infra/parser"

// NewXLSXParser injects all dependencies in any object that implements parser.Parser and return it
func NewXLSXParser() parser.Parser {
	return parser.Factory(parser.Default)
}

// NewCSVParser injects all dependencies in any object that implements parser.Parser and return it
func NewCSVParser() parser.Parser {
	return parser.Factory(parser.CSV)
}
