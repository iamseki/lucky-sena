package generator

// Generator expose the methods to be implemented by who knows how to do it
type Generator interface {
	Generate(o Options) []GenaretedBet
}

// Options to used in Generated method
type Options struct {
	BetsToGenerate  int
	ExcludedNumbers []int
}

// GenaretedBet is the properties of bet generated
type GenaretedBet struct {
	Numbers []int
}
