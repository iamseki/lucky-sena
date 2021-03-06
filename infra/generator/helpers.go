package generator

import (
	"strconv"
	"strings"
)

func containValue(numbers []int, value int) bool {
	for _, n := range numbers {
		if n == value {
			return true
		}
	}
	return false
}

// ConvertCSVIntoIntSlice covert some string separeted by commas into an int slice
func ConvertCSVIntoIntSlice(csvString string) []int {
	separetedByComma := strings.Split(csvString, ",")
	var intSlice []int
	for _, value := range separetedByComma {
		n, _ := strconv.Atoi(value)
		intSlice = append(intSlice, n)
	}
	return intSlice
}
