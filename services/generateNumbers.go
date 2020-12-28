package services

import (
	"log"
	"lucky-sena/models"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// BetGenerator service that implements IFactory
type BetGenerator struct {
	bets            int
	excludedNumbers []int
}

func newBetGenerator() *BetGenerator {
	return &BetGenerator{}
}

// SetNumberOfBets sets the number of generated bets
func (b *BetGenerator) SetNumberOfBets(bets int) {
	b.bets = bets
}

// SetExcludedNumbers csv format of the numbers to ignore in generator
func (b *BetGenerator) SetExcludedNumbers(excludedNumbers string) {
	separetedByComma := strings.Split(excludedNumbers, ",")
	for _, value := range separetedByComma {
		n, _ := strconv.Atoi(value)
		b.excludedNumbers = append(b.excludedNumbers, n)
	}
	log.Printf("Excluded numbers: %v\n", b.excludedNumbers)
}

// Run the Number Generator code, printing every generated bet
func (b *BetGenerator) Run() {
	var wg sync.WaitGroup
	var bet models.Bet
	const price float32 = 4.5

	bet.Date = time.Now()
	s := rand.NewSource(time.Now().UnixNano())
	random := rand.New(s)

	for c := 0; c < b.bets; c++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			numbers := make([]int, 6)
			generateNumbers(numbers, random, b.excludedNumbers)
			log.Printf("bet(%d) - %v\n", c, numbers)
			bet.Numbers = append(bet.Numbers, models.Balls{Bet: numbers})
		}(c)
	}
	wg.Wait()

	bet.Coast = price * float32(len(bet.Numbers))
	log.Println(bet)
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

func containValue(numbers []int, value int) bool {
	for _, n := range numbers {
		if n == value {
			return true
		}
	}
	return false
}
