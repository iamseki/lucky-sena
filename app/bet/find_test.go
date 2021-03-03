package betusecases

import (
	"lucky-sena/domain"
	"testing"
)

type mockFindError struct {
	message string
}

func (e *mockFindError) Error() string {
	return e.message
}

func TestFind(t *testing.T) {
	r := newFakeFindRepository()
	sut := NewFindBet(r)
	bets, _ := sut.Find()
	if len(bets) != 2 {
		t.Error("expected find 2 bets but got", len(bets))
	}
}

func TestFindError(t *testing.T) {
	r := newFakeFindRepository()
	r.FindMock = func() ([]domain.Bet, error) {
		return []domain.Bet{}, &mockFindError{message: "Error trying to Find Bets"}
	}
	sut := NewFindBet(r)
	_, err := sut.Find()
	if err == nil {
		t.Error("expected err to be not nil but got", err)
	}
}
