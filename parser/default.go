package parser

import (
	"strconv"
	"strings"
)

type defaultParser struct{}

func newDefaultParser() Parser {
	return &defaultParser{}
}

func (parser *defaultParser) Parse(excludedNumbers string) []int {
	separetedByComma := strings.Split(excludedNumbers, ",")
	var parsedExludedNumbers []int
	for _, value := range separetedByComma {
		n, _ := strconv.Atoi(value)
		parsedExludedNumbers = append(parsedExludedNumbers, n)
	}
	return parsedExludedNumbers
}
