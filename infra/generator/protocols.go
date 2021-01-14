package generator

type Generator interface {
	Generate(o Options) []GenaretedBet
}

type Options struct {
	BetsToGenerate  int
	ExcludedNumbers []int
}

type GenaretedBet struct {
	Numbers []int
}
