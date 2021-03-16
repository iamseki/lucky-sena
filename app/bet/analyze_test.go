package betusecases_test

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
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
