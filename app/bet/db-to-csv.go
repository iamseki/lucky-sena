package betusecases

import "lucky-sena/app/protocols"

type DbBetToCsv struct {
	repository protocols.FindBetRepository
	betToCsv   protocols.CSVConverter
}

func NewDbBetToCsv(r protocols.FindBetRepository, c protocols.CSVConverter) *DbBetToCsv {
	return &DbBetToCsv{
		repository: r,
		betToCsv:   c,
	}
}

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
