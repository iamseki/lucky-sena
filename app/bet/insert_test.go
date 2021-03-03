package betusecases

import (
	"lucky-sena/domain"
	"testing"
)

type mockInsertError struct {
	message string
}

func (m *mockInsertError) Error() string {
	return m.message
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

func TestInsertBets(t *testing.T) {
	r := newInsertFakeRepository()
	sut := NewInsertBets(r)

	err := sut.InsertBets(makeMockedBets())
	if err != nil {
		t.Error("expect error to be nil but got", err)
	}
}

func TestInsertBetsFails(t *testing.T) {
	r := newInsertFakeRepository()
	r.InsertManyMock = func(b []domain.Bet) error { return &mockInsertError{message: "Error on InsertMany"} }
	sut := NewInsertBets(r)

	err := sut.InsertBets(makeMockedBets())
	if err == nil {
		t.Error("expect error to be not nil but got", err)
	}
}
