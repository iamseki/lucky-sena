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
	s := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s)
	var wg sync.WaitGroup
	var bets []GenaretedBet

	for c := 0; c < options.BetsToGenerate-1; c++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			numbers := make([]int, 6)
			generateNumbers(numbers, random, options.ExcludedNumbers)
			bets = append(bets, GenaretedBet{numbers})
		}(c)
	}
	wg.Wait()

	bets = includeMergedBet(bets, options.BetsToGenerate)

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

func includeMergedBet(bets []GenaretedBet, betsMaded int) []GenaretedBet {
	m := make(map[int]int)

	for _, bet := range bets {
		for _, number := range bet.Numbers {
			if _, exists := m[number]; !exists {
				m[number] = 1
			} else {
				m[number] += 1
			}
		}
	}

	merged := GenaretedBet{}

	// a number can appeared at max the numbers of bets maded
	// cause there's no repeating numbers in a single bet
	maxBetsCount := betsMaded

	for maxBetsCount >= 1 {
		for number, appearedTimes := range m {
			if len(merged.Numbers) == 6 {
				break
				// ranking by times appeared in bets
			} else if appearedTimes == maxBetsCount {
				merged.Numbers = append(merged.Numbers, number)
			}
		}

		maxBetsCount--
	}

	sort.Ints(merged.Numbers)

	bets = append(bets, merged)
	return bets
}
