package scrapper

import (
	"log"
	"lucky-sena/domain"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

type CollyScrapper struct{}

func NewCollyScrapper() *CollyScrapper {
	return &CollyScrapper{}
}

func (s *CollyScrapper) Scrap(endpoint string) (domain.Bet, error) {
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
	err := c.Visit(endpoint)
	if err != nil {
		log.Fatalln(err)
	}

	wg.Wait()

	b := domain.Bet{Numbers: numbers, Code: gameCode, Date: date}
	return b, nil
}
