package betusecases

import "lucky-sena/app/protocols"

type DbBetToCsv struct {
	Repository protocols.FindBetRepository
	BetToCsv   protocols.CSVConverter
}

func (usecase *DbBetToCsv) ConvertBetsToCsv() error {
	bets, err := usecase.Repository.Find()
	if err != nil {
		return err
	}
	err = usecase.BetToCsv.Convert(bets)
	if err != nil {
		return err
	}
	return nil
}
