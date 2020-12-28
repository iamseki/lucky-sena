package services

import (
	"math/rand"
	"testing"
	"time"
)

func TestGenerateNumbers(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	numbers := make([]int, 6)
	excludedNumbers := []int{0}

	generateNumbers(numbers, r, excludedNumbers)

	for idx, n := range numbers {
		if n <= 0 || n > 60 {
			t.Errorf("generateNumbers() failed, expected any number between 1 and 60, got: %v at index: %v", n, idx)
		}
	}
}

func TestGenerateNumbersWithExcludes(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	numbers := make([]int, 6)
	excludedNumbers := []int{1, 2, 3, 4, 5, 6, 10, 60}

	generateNumbers(numbers, r, excludedNumbers)
	for idx, n := range numbers {
		if containValue(excludedNumbers, n) {
			t.Errorf(`generateNumbers() failed, expected any number between 1 and 60 AND not in: %v 
got: %v at index: %v`, excludedNumbers, n, idx)
		}
	}
}
