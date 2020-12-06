package main

import (
	"flag"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var bets int
	var lastResult string

	flag.IntVar(&bets, "b", 7, "b is equal the number of bets to be done")
	flag.StringVar(&lastResult, "e", "", "e is the numbers to exclude in csv format: 1,2,3,4,5,6")
	flag.Parse()

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	split := strings.Split(lastResult, ",")
	lastResults := make([]int, len(split))
	for idx, value := range split {
		lastResults[idx], _ = strconv.Atoi(value)
	}

	log.Printf("Last Result: %v\n", lastResults)

	for c := 0; c < bets; c++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()

			numbers := make([]int, 6)
			generateNumbers(numbers, r, lastResults)
			log.Printf("bet(%d) - %v\n", c, numbers)
		}(c)
	}

	wg.Wait()
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
}

func contains(s []int, v int) bool {
	for _, n := range s {
		if n == v {
			return true
		}
	}
	return false
}
