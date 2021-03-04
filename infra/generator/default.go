package generator

import (
	"math/rand"
	"sort"
	"sync"
	"time"
)

type defaultGenerator struct{}

func newDefaultGenerator() Generator {
	return &defaultGenerator{}
}

func (generator *defaultGenerator) Generate(options Options) []GenaretedBet {
	var wg sync.WaitGroup
	var bets []GenaretedBet

	for c := 0; c < options.BetsToGenerate; c++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			numbers := make([]int, 6)
			s := rand.NewSource(time.Now().UnixNano())
			random := rand.New(s)
			generateNumbers(numbers, random, options.ExcludedNumbers)
			bets = append(bets, GenaretedBet{numbers})
		}(c)
	}
	wg.Wait()

	return bets
}

func generateNumbers(numbers []int, random *rand.Rand, lastResults []int) {
	for i := 0; i < 6; i++ {
		randomNumber := random.Intn(61)
		if containValue(lastResults, randomNumber) || containValue(numbers, randomNumber) {
			i--
			continue
		}
		numbers[i] = randomNumber
	}
	sort.Ints(numbers)
}
