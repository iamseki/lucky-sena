package generator

type Generator interface {
	Generate(o Options) []Bet
}

type Options struct {
	Bets            int
	ExcludedNumbers []int
}

type Bet struct {
	Numbers []int
}
