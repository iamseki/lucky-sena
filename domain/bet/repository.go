package bet

type AddBetRepository interface {
	Add(Bet) error
}

type FindBetRepository interface {
	Find() ([]Bet, error)
}

type FindBetByCodeRepository interface {
	FindBetByCode(code int) (Bet, error)
}
