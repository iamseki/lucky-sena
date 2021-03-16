package betusecases_test

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"testing"
	"time"
)

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
func TestAddBetSuccess(t *testing.T) {
	r := newFakeAddRepository()
	a := betusecases.NewAddBet(r)

	bet, _ := a.AddBet(domain.Bet{
		Numbers: []int{1, 2, 3, 4, 5, 6},
		Code:    17,
		Date:    time.Now(),
	})

	if bet.ID != "any-id" {
		t.Error("expect bet.ID to be any-id but got", bet.ID)
	}
}

func TestAddBetFails(t *testing.T) {
	r := newFakeAddRepository()
	r.AddFn = func(b domain.Bet) (domain.BetModel, error) {
		return domain.BetModel{}, newTestError("Error on Add Bet at repository")
	}
	a := betusecases.NewAddBet(r)

	_, err := a.AddBet(domain.Bet{
		Numbers: []int{1, 2, 3, 4, 5, 6},
		Code:    17,
		Date:    time.Now(),
	})

	if err == nil {
		t.Error("err must not be empty but got", err)
	}
}
