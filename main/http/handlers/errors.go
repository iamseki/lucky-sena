package handlers

type generateBetError struct {
	message string
}

func (g generateBetError) Error() string {
	return g.message
}
