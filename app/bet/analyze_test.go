package betusecases

import (
	"lucky-sena/app/protocols"
	"lucky-sena/domain"
	"reflect"
	"testing"
	"time"
)

func makeMockedBets() []domain.Bet {
	return []domain.Bet{
		{Numbers: []int{1, 2, 3, 4, 5, 6}, Code: 1, Date: time.Now()},
		{Numbers: []int{1, 2, 3, 4, 5, 6}, Code: 2, Date: time.Now()},
	}
}

func makeSut() *AnalyzeBet {
	r := newFakeAnalyzeRepository()
	return NewAnalyzeBet(r)
}

func TestNextBetCode(t *testing.T) {
	sut := makeSut()

	nextGame := sut.NextBetCode()

	if nextGame != 3 {
		t.Error("expect bet to be 3 but got", nextGame)
	}
}

func TestIsBetAlreadyWon(t *testing.T) {
	sut := makeSut()

	if isBetWon := sut.IsBetAlreadyWon([]int{1, 2, 3, 4, 5, 6}); !isBetWon {
		t.Error("expected bet to have already wan but got", isBetWon)
	}

	if isBetWon := sut.IsBetAlreadyWon([]int{1, 2, 3, 4, 5, 7}); isBetWon {
		t.Error("expected bet to not have won but got ", isBetWon)
	}
}

type fakeRepository struct{}

func (f *fakeRepository) Find() ([]domain.Bet, error) {
	return makeMockedBets(), nil
}

func (f *fakeRepository) FindBetByCode(code int) (domain.Bet, error) {
	bets := makeMockedBets()
	for _, b := range bets {
		if code == b.Code {
			return b, nil
		}
	}
	return domain.Bet{}, nil
}

func (f *fakeRepository) FindBetsByNumbers(numbers []int) ([]domain.Bet, error) {
	fakeBets := makeMockedBets()
	const firstBet = 0
	if reflect.DeepEqual(numbers, fakeBets[firstBet].Numbers) {
		return fakeBets, nil
	}
	return nil, nil
}

func newFakeAnalyzeRepository() protocols.AnalyzeBetRepository {
	return &fakeRepository{}
}
