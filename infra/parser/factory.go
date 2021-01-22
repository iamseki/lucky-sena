package parser

type ParserType string

const (
	Default      ParserType = "default"
	XlsxExcelize            = "xlsxExcelize"
)

func Factory(pt ParserType) Parser {
	switch pt {
	case XlsxExcelize:
		return newXlsxExcelizeParser()
	default:
		return newXlsxExcelizeParser()
	}
}
