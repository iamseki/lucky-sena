package converter

import (
	"encoding/csv"
	"fmt"
	"log"
	"lucky-sena/domain"
	"os"
	"strconv"
	"strings"
	"time"
)

// BetToCsv represents an struct who knows how to Convert Bets to Csv
type BetToCsv struct {
	PrefixFilename string
}

// NewBetToCsvConverter returns an instance of BetToCsv struct
func NewBetToCsvConverter(filename string) *BetToCsv {
	return &BetToCsv{
		PrefixFilename: filename,
	}
}

// Convert converts all bets in csv format
func (bc *BetToCsv) Convert(bets []domain.Bet) error {
	convertDate := time.Now().Format("2006-01-02")
	filename := bc.PrefixFilename + "-csv-" + convertDate + ".csv"

	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln(err)
	}

	records := [][]string{
		{"numbers", "code", "date"},
	}

	for _, b := range bets {
		code := strconv.Itoa(b.Code)
		date := b.Date.Format("2006-01-02")
		numbers := intArrayToString(b.Numbers, "|")
		records = append(records, []string{numbers, code, date})
	}

	w := csv.NewWriter(file)
	err = w.WriteAll(records)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func intArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
