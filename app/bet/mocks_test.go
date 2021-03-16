package betusecases_test

type testError struct {
	message string
}

func (te *testError) Error() string {
	return te.message
}

func newTestError(message string) *testError {
	return &testError{message}
}
