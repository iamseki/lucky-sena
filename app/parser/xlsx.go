package parser

import (
	"log"
	"lucky-sena/domain/bet"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// XlsxParser struct that implements IFactory
type XlsxParser struct{}

func newXlsxParser() *XlsxParser {
	return &XlsxParser{}
}

func (xw *XlsxParser) Parse(options Options) []bet.Bet {
	file, err := excelize.OpenFile(options.FileName)
	if err != nil {
		log.Fatalln(err)
	}

	xlsxSheet := file.GetSheetName(file.GetActiveSheetIndex())
	rows := file.GetRows(xlsxSheet)
	var bets []bet.Bet

	for _, row := range rows {
		betString := row[2:]
		betInt := make([]int, 6)

		if _, err := strconv.Atoi(betString[0]); err != nil {
			continue
		}
		for i, stringNumber := range betString {
			numInt, err := strconv.Atoi(stringNumber)
			if err != nil {
				break
			}
			betInt[i] = numInt
		}
		sort.Ints(betInt)

		gameCode := 0
		code, _ := strconv.Atoi(row[gameCode])

		gameDate := 1
		splited := strings.Split(row[gameDate], "/")
		var DAY, MONTH, YEAR = 0, 1, 2
		standerized := splited[YEAR] + splited[MONTH] + splited[DAY]
		standardSampleLayout := "20060102"

		date, err := time.Parse(standardSampleLayout, standerized)
		if err != nil {
			log.Fatalln(err)
		}

		bets = append(bets, bet.Bet{betInt, code, date, bet.Coast})
	}

	return bets
}
