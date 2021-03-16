package betusecases_test

import (
	"lucky-sena/domain"
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
