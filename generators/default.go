package generators

import (
	"log"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type defaultGenerator struct{}

func newDefaultGenerator() Generator {
	return &defaultGenerator{}
}

func (generator *defaultGenerator) Generate(options Options) []Bet {
	s := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s)
	var wg sync.WaitGroup
	var bets []Bet

	for c := 0; c < options.Bets; c++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			numbers := make([]int, 6)
			generateNumbers(numbers, random, options.ExcludedNumbers)
			log.Printf("bet(%d) - %v\n", c, numbers)
			bets = append(bets, Bet{Numbers: numbers})
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
