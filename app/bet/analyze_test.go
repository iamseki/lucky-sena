package betusecases_test

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"reflect"
	"testing"
)

func TestNextBetCode(t *testing.T) {
	r := newFakeAnalyzeRepository()
	sut := betusecases.NewAnalyzeBet(r)

	nextGame, _ := sut.NextBetCode()

	if nextGame != 3 {
		t.Error("expect bet to be 3 but got", nextGame)
	}
}

func TestNextBetCodeFails(t *testing.T) {
	r := newFakeAnalyzeRepository()
	r.FindMock = func() ([]domain.Bet, error) {
		return []domain.Bet{}, newTestError("Error on analyze repository")
	}
	sut := betusecases.NewAnalyzeBet(r)

	_, err := sut.NextBetCode()

	if err == nil {
		t.Error("expect err to be nil but got", err)
	}
}

func TestIsBetAlreadyWon(t *testing.T) {
	r := newFakeAnalyzeRepository()
	sut := betusecases.NewAnalyzeBet(r)

	if isBetWon, _ := sut.IsBetAlreadyWon([]int{1, 2, 3, 4, 5, 6}); !isBetWon {
		t.Error("expected bet to have already wan but got", isBetWon)
	}

	if isBetWon, _ := sut.IsBetAlreadyWon([]int{1, 2, 3, 4, 5, 7}); isBetWon {
		t.Error("expected bet to not have won but got ", isBetWon)
	}
}

func TestIsBetAlreadyWonFails(t *testing.T) {
	r := newFakeAnalyzeRepository()
	r.FindBetByNumbersMock = func(numbers []int) ([]domain.Bet, error) {
		return []domain.Bet{}, newTestError("Error on analyze repository")
	}
	sut := betusecases.NewAnalyzeBet(r)

	if _, err := sut.IsBetAlreadyWon([]int{1, 2, 3, 4, 5, 6}); err == nil {
		t.Error("expect err to not be nil, got: ", err)
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
