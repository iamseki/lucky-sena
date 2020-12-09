package services

import (
	"fmt"
	"log"
	"lucky-sena/models"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type XlsxWriter struct {
}

func newXlsxWriter() *XlsxWriter {
	return &XlsxWriter{}
}

func (xw *XlsxWriter) Run() {
	f, err := excelize.OpenFile("mega_sena_asloterias_ate_concurso_2325_sorteio.xlsx")
	if err != nil {
		log.Fatalln(err)
	}

	sheet := f.GetSheetName(f.GetActiveSheetIndex())

	rows := f.GetRows(sheet)

	for code, row := range rows {
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

		splited := strings.Split(row[1], "/")
		// YYYY/MM/DD
		standerized := splited[2] + splited[1] + splited[0]
		date, err := time.Parse("20060102", standerized)
		if err != nil {
			fmt.Println(err)
		}

		//log.Printf("%v - [%v] - Bet: %v", row[1], code, betInt)
		result := models.Result{Code: code, Date: date, Bet: betInt}

		// PUT the result in mongo results collection
		log.Println(result)
	}
}
