package generators

type GeneratorType string

// types available of Generators
const (
	Default   GeneratorType = "default"
	CustomBar               = "customBar"
)

func MakeGenerator(t GeneratorType) Generator {
	switch t {
	case Default:
		return newDefaultGenerator()
	case CustomBar:
		return newDefaultGenerator()
	default:
		return newDefaultGenerator()
	}
}
