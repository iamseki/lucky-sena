package services

import (
	"fmt"
	"log"
	"lucky-sena/database"
	"lucky-sena/models"
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

	// Connect to mongodb database
	m := database.Factory(database.MongoDB).(*database.Mongo)
	defer m.Client.Disconnect(m.Ctx)

	resultsCollection := m.Client.Database("sena").Collection("results")
	// Read xlsx
	f, err := excelize.OpenFile(xw.fileName)
	if err != nil {
		log.Fatalln(err)
	}

	sheet := f.GetSheetName(f.GetActiveSheetIndex())
	rows := f.GetRows(sheet)

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

		// Code its in first position of the row array
		code, _ := strconv.Atoi(row[0])
		// row[1] include the date in the format DD/MM/YYYY
		splited := strings.Split(row[1], "/")
		// from DD/MM/YYYY to  YYYYMMDD
		standerized := splited[2] + splited[1] + splited[0]
		// 20060102 is the standard layout string expected
		date, err := time.Parse("20060102", standerized)
		if err != nil {
			fmt.Println(err)
		}

		//log.Printf("%v - [%v] - Bet: %v", row[1], code, betInt)
		//results = append(results, models.Result{Code: code, Date: date, Bet: betInt})
		wg.Add(1)
		go func(code int, date time.Time, betInt []int) {
			result := &models.Result{Code: code, Date: date, Bet: betInt}
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
