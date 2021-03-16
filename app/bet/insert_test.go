package betusecases_test

import (
	betusecases "lucky-sena/app/bet"
	"lucky-sena/domain"
	"testing"
)

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

func TestInsertBets(t *testing.T) {
	r := newInsertFakeRepository()
	sut := betusecases.NewInsertBets(r)

	err := sut.InsertBets(makeMockedBets())
	if err != nil {
		t.Error("expect error to be nil but got", err)
	}
}

func TestInsertBetsFails(t *testing.T) {
	r := newInsertFakeRepository()
	r.InsertManyMock = func(b []domain.Bet) error { return newTestError("Error on InsertMany") }
	sut := betusecases.NewInsertBets(r)

	err := sut.InsertBets(makeMockedBets())
	if err == nil {
		t.Error("expect error to be not nil but got", err)
	}
}
