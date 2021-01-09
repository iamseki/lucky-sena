package parser

type Parser interface {
	Parse(data string) []int
}
