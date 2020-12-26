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

// SetBet sets the number of generated bets
func (b *BetGenerator) SetBet(bets int) {
	b.bets = bets
}

// SetExcludedNumbers csv format of the numbers to ignore in generator
func (b *BetGenerator) SetExcludedNumbers(excludeArgs string) {
	split := strings.Split(excludeArgs, ",")

	for _, value := range split {
		n, _ := strconv.Atoi(value)
		b.excludedNumbers = append(b.excludedNumbers, n)
	}

	log.Printf("Excluded numbers: %v\n", b.excludedNumbers)
}

// Run the runnable code of BetGenerator service
func (b *BetGenerator) Run() {
	var wg sync.WaitGroup
	var bet models.Bet

	const price float32 = 4.5

	bet.Date = time.Now()

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for c := 0; c < b.bets; c++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()

			numbers := make([]int, 6)
			generateNumbers(numbers, r, b.excludedNumbers)
			log.Printf("bet(%d) - %v\n", c, numbers)
			bet.Numbers = append(bet.Numbers, models.Balls{Bet: numbers})
		}(c)
	}

	wg.Wait()

	bet.Coast = price * float32(len(bet.Numbers))
	// PUT TE BET DOCUMENT INSIDE MONGODB
	log.Println(bet)
}

func generateNumbers(numbers []int, r *rand.Rand, lastResults []int) {
	for i := 0; i < 6; i++ {
		rn := r.Intn(61)

		if contains(lastResults, rn) || contains(numbers, rn) {
			i--
			continue
		}

		numbers[i] = rn
	}

	sort.Ints(numbers)
}

func contains(s []int, v int) bool {
	for _, n := range s {
		if n == v {
			return true
		}
	}
	return false
}
