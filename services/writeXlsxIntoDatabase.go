package services

import (
	"fmt"
	"log"
	"lucky-sena/database"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// XlsxWriter struct that implements IFactory
type XlsxWriter struct {
	fileName string
}

func newXlsxWriter() *XlsxWriter {
	return &XlsxWriter{}
}

// SetFileName sets the name of the file to read
func (xw *XlsxWriter) SetFileName(name string) {
	xw.fileName = name
}

// Run method implemented by XlsxWriter struct
func (xw *XlsxWriter) Run() {
	var wg sync.WaitGroup

	m := database.Factory(database.MongoDB).(*database.Mongo)
	defer m.Client.Disconnect(m.Ctx)

	resultsCollection := m.Client.Database("sena").Collection("results")

	file, err := excelize.OpenFile(xw.fileName)
	if err != nil {
		log.Fatalln(err)
	}

	sheet := file.GetSheetName(file.GetActiveSheetIndex())
	rows := file.GetRows(sheet)

	for _, row := range rows {
		betStr := row[2:]
		betInt := make([]int, 6)

		if _, err := strconv.Atoi(betStr[0]); err != nil {
			continue
		}
		for i, numStr := range betStr {
			numInt, err := strconv.Atoi(numStr)
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
		standardSampleLayout := "gameDate"
		date, err := time.Parse(standardSampleLayout, standerized)
		if err != nil {
			fmt.Println(err)
		}

		//log.Printf("%v - [%v] - Bet: %v", row[1], code, betInt)
		//results = append(results, models.Result{Code: code, Date: date, Bet: betInt})
		wg.Add(1)
		go func(code int, date time.Time, betInt []int) {
			result := &Result{Code: code, Date: date, Bet: betInt}
			_, err = resultsCollection.InsertOne(m.Ctx, result)
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}(code, date, betInt)
	}

	wg.Wait()

	log.Println("Results was inserted into results collection in sena database !")
}
