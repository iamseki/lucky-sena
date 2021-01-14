package parser

type ParserType string

const (
	Default ParserType = "default"
	Xlsx               = "xlsx"
)

func Factory(pt ParserType) Parser {
	switch pt {
	case Xlsx:
		return newXlsxParser()
	default:
		return newXlsxParser()
	}
}
