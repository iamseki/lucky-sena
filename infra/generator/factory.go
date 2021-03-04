package generator

// Type is the type of instance to be used to generate numbers
type Type string

// types available of Generators
const (
	Default   Type = "default"
	CustomBar      = "customBar"
)

// Factory returns an instance that knows how to generate numbers accordlying to t Type
func Factory(t Type) Generator {
	switch t {
	case Default:
		return newDefaultGenerator()
	case CustomBar:
		return newDefaultGenerator()
	default:
		return newDefaultGenerator()
	}
}
