package scrapper

import (
	"log"
	"strconv"
	"strings"
	"time"
)

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
