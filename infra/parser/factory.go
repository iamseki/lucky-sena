package parser

type ParserType string

const (
	Default      ParserType = "default"
	XlsxExcelize            = "xlsxExcelize"
	CSV                     = "csv"
)

func Factory(pt ParserType) Parser {
	switch pt {
	case XlsxExcelize:
		return newXlsxExcelizeParser()
	case CSV:
		return newCsvParser()
	default:
		return newXlsxExcelizeParser()
	}
}
