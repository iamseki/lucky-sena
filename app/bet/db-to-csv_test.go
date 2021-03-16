package betusecases_test

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"testing"
)

func TestConvertBetsToCsv(t *testing.T) {
	r := newFakeFindRepository()
	c := newFakeCSVConverter()
	sut := betusecases.NewDbBetToCsv(r, c)

	err := sut.ConvertBetsToCsv()
	if err != nil {
		t.Error("err must be nil but got", err)
	}
}

func TestConvertBetsToCsvFindFails(t *testing.T) {
	r := newFakeFindRepository()
	r.FindMock = func() ([]domain.Bet, error) {
		return []domain.Bet{}, newTestError("Error on DbBetToCsv usecase in Find method")
	}
	c := newFakeCSVConverter()
	sut := betusecases.NewDbBetToCsv(r, c)

	err := sut.ConvertBetsToCsv()
	if err == nil {
		t.Error("err must not be nil but got", err)
	}
}

func TestConvertBetsToCsvConverterFails(t *testing.T) {
	r := newFakeFindRepository()
	c := newFakeCSVConverter()
	c.ConvertMock = func(b []domain.Bet) error {
		return newTestError("Error on DbBetToCsv usecase in Converter method")
	}
	sut := betusecases.NewDbBetToCsv(r, c)

	err := sut.ConvertBetsToCsv()
	if err == nil {
		t.Error("err must not be nil but got", err)
	}
}
