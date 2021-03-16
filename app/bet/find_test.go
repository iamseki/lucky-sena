package betusecases_test

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"testing"
)

func TestFind(t *testing.T) {
	r := newFakeFindRepository()
	sut := betusecases.NewFindBet(r)
	bets, _ := sut.Find()
	if len(bets) != 2 {
		t.Error("expected find 2 bets but got", len(bets))
	}
}

func TestFindError(t *testing.T) {
	r := newFakeFindRepository()
	r.FindMock = func() ([]domain.Bet, error) {
		return []domain.Bet{}, newTestError("Error trying to Find Bets")
	}
	sut := betusecases.NewFindBet(r)
	_, err := sut.Find()
	if err == nil {
		t.Error("expected err to be not nil but got", err)
	}
}
