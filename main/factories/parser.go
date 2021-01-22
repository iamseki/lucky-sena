package factories

import "lucky-sena/infra/parser"

func NewXLSXParser() parser.Parser {
	return parser.Factory(parser.Default)
}
