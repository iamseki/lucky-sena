package betusecases_test

import (
	"lucky-sena/domain"
	"reflect"
	"time"
)

type testError struct {
	message string
}

func (te *testError) Error() string {
	return te.message
}

func newTestError(message string) *testError {
	return &testError{message}
}

func makeMockedBets() []domain.Bet {
	return []domain.Bet{
		{Numbers: []int{1, 2, 3, 4, 5, 6}, Code: 1, Date: time.Now()},
		{Numbers: []int{1, 2, 3, 4, 5, 6}, Code: 2, Date: time.Now()},
	}
}

type fakeRepository struct {
	FindMock             func() ([]domain.Bet, error)
	FindBetByCodeMock    func(code int) (domain.Bet, error)
	FindBetByNumbersMock func(numbers []int) ([]domain.Bet, error)
}

func (f *fakeRepository) Find() ([]domain.Bet, error) {
	return f.FindMock()
}

func (f *fakeRepository) FindBetByCode(code int) (domain.Bet, error) {
	return f.FindBetByCodeMock(code)
}

func (f *fakeRepository) FindBetsByNumbers(numbers []int) ([]domain.Bet, error) {
	return f.FindBetByNumbersMock(numbers)
}

func newFakeAnalyzeRepository() *fakeRepository {
	return &fakeRepository{
		FindMock: func() ([]domain.Bet, error) {
			return makeMockedBets(), nil
		},
		FindBetByCodeMock: func(code int) (domain.Bet, error) {
			bets := makeMockedBets()
			for _, b := range bets {
				if code == b.Code {
					return b, nil
				}
			}
			return domain.Bet{}, nil
		},
		FindBetByNumbersMock: func(numbers []int) ([]domain.Bet, error) {
			fakeBets := makeMockedBets()
			const firstBet = 0
			if reflect.DeepEqual(numbers, fakeBets[firstBet].Numbers) {
				return fakeBets, nil
			}
			return nil, nil
		},
	}
}

type fakeAddRepository struct {
	AddFn func(b domain.Bet) (domain.BetModel, error)
}

func newFakeAddRepository() *fakeAddRepository {
	return &fakeAddRepository{
		AddFn: func(b domain.Bet) (domain.BetModel, error) {
			return domain.BetModel{
				ID:      "any-id",
				Numbers: b.Numbers,
				Code:    b.Code,
				Date:    b.Date,
			}, nil
		},
	}
}

func (fr *fakeAddRepository) Add(b domain.Bet) (domain.BetModel, error) {
	return fr.AddFn(b)
}

type fakeFindRepository struct {
	FindMock func() ([]domain.Bet, error)
}

func (fr *fakeFindRepository) Find() ([]domain.Bet, error) {
	return fr.FindMock()
}

func newFakeFindRepository() *fakeFindRepository {
	return &fakeFindRepository{
		FindMock: func() ([]domain.Bet, error) {
			return makeMockedBets(), nil
		},
	}
}

type fakeCSVConverter struct {
	ConvertMock func([]domain.Bet) error
}

func (fc *fakeCSVConverter) Convert(bets []domain.Bet) error {
	return fc.ConvertMock(bets)
}

func newFakeCSVConverter() *fakeCSVConverter {
	return &fakeCSVConverter{
		ConvertMock: func(b []domain.Bet) error { return nil },
	}
}

type fakeInserRepositoy struct {
	InsertManyMock func([]domain.Bet) error
}

func (fr *fakeInserRepositoy) InsertMany(bets []domain.Bet) error {
	return fr.InsertManyMock(bets)
}

func newInsertFakeRepository() *fakeInserRepositoy {
	return &fakeInserRepositoy{
		InsertManyMock: func(b []domain.Bet) error {
			return nil
		},
	}
}
