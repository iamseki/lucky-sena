package betusecases

import "lucky-sena/app/protocols"

// DbBetToCsv represents a struct that implements domain.BetCsvConverter interface
type DbBetToCsv struct {
	repository protocols.FindBetRepository
	betToCsv   protocols.CSVConverter
}

// NewDbBetToCsv returns an instance of DbBetToCsv struct
func NewDbBetToCsv(r protocols.FindBetRepository, c protocols.CSVConverter) *DbBetToCsv {
	return &DbBetToCsv{
		repository: r,
		betToCsv:   c,
	}
}

// ConvertBetsToCsv converts all bets in the database to a csv format
func (usecase *DbBetToCsv) ConvertBetsToCsv() error {
	bets, err := usecase.repository.Find()
	if err != nil {
		return err
	}
	err = usecase.betToCsv.Convert(bets)
	if err != nil {
		return err
	}
	return nil
}
