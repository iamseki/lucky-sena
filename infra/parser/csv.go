package parser

import (
	"encoding/csv"
	"io"
	"log"
	"lucky-sena/domain"
	"os"
	"strconv"
	"strings"
	"time"
)

type CsvBetParser struct{}

func newCsvParser() *CsvBetParser {
	return &CsvBetParser{}
}

func (p *CsvBetParser) Parse(options Options) []domain.Bet {
	var bets []domain.Bet
	standardSampleLayout := "2006-01-02"
	NUMBERS_INDEX, CODE_INDEX, DATE_INDEX := 0, 1, 2

	r, os := getCSVReader(options.FileName)
	defer os.Close()

	header := getCSVNextRecord(r)
	log.Println(header)
	for {
		record := getCSVNextRecord(r)
		if record == nil {
			break
		}

		date := parseDate(record[DATE_INDEX], standardSampleLayout)
		code := parseInt(record[CODE_INDEX])
		numbers := parseNumbers(record[NUMBERS_INDEX])
		bets = append(bets, domain.Bet{Code: code, Date: date, Numbers: numbers})
	}

	return bets
}

func getCSVReader(filename string) (*csv.Reader, *os.File) {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	return r, csvfile
}

func getCSVNextRecord(r *csv.Reader) []string {
	record, err := r.Read()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}
	return record
}
func parseDate(input string, standardSampleLayout string) time.Time {
	date, err := time.Parse(standardSampleLayout, input)
	if err != nil {
		log.Fatalln(err)
	}
	return date
}

func parseInt(input string) int {
	n, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalln(err)
	}
	return n
}

func parseNumbers(csvString string) []int {
	separetedByComma := strings.Split(csvString, "|")
	var intSlice []int
	for _, value := range separetedByComma {
		n, _ := strconv.Atoi(value)
		intSlice = append(intSlice, n)
	}
	return intSlice
}
