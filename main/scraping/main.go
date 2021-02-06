package main

import (
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	var gameCode int
	var numbers []int
	var date time.Time
	var wg sync.WaitGroup
	wg.Add(2)
	// On every element call callback
	c.OnHTML("body", func(e *colly.HTMLElement) {
		concatenedStringTextBody := e.Text
		log.Println("text scraped from body:", concatenedStringTextBody)
		gameCode = getGameCodeFromString(concatenedStringTextBody)
		date = getDateWinnerFromString(concatenedStringTextBody)
	})

	c.OnHTML(".num_sorteio", func(e *colly.HTMLElement) {
		defer wg.Done()
		concatenedStringBet := e.Text
		log.Println("text scraped ul containing the winner bet:", concatenedStringBet)
		numbers = getNumbersFromString(concatenedStringBet)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Scraping bets from", r.URL.String())
	})

	// Start scraping
	err := c.Visit("http://www1.caixa.gov.br/loterias/loterias/megasena/megasena_pesquisa_new.asp")
	if err != nil {
		log.Fatalln(err)
	}

	wg.Wait()

	b := Bet{numbers: numbers, game: gameCode, date: date}
	log.Println("Winner bet:", b)
}

type Bet struct {
	numbers []int
	game    int
	date    time.Time
}

func getNumbersFromString(s string) []int {
	var numbers []int
	for i := 0; i < len(s); i += 2 {
		num, err := strconv.Atoi(s[i : i+2])
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func getGameCodeFromString(s string) int {
	codeString := s[:strings.Index(s, "|")]
	gameCode, err := strconv.Atoi(codeString)
	if err != nil {
		log.Fatalln(err)
	}
	return gameCode
}

func getDateWinnerFromString(s string) time.Time {
	standardSampleLayout := "02/01/2006"
	firstDashPos := strings.Index(s, "/")
	stringDate := s[firstDashPos-2 : firstDashPos+8]

	date, err := time.Parse(standardSampleLayout, stringDate)
	if err != nil {
		log.Fatalln(err)
	}
	return date
}
