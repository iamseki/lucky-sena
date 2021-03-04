package parser

// Type is the type of instance to be used to generate numbers
type Type string

// defaults types that can be used to inject the parser
const (
	Default      Type = "default"
	XlsxExcelize      = "xlsxExcelize"
	CSV               = "csv"
)

// Factory returns an instance that knows how to parse csv accordlying to t Type
func Factory(pt Type) Parser {
	switch pt {
	case XlsxExcelize:
		return newXlsxExcelizeParser()
	case CSV:
		return newCsvParser()
	default:
		return newXlsxExcelizeParser()
	}
}
