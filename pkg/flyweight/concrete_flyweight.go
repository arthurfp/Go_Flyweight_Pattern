package flyweight

// ConcreteFlyweight is a concrete implementation of the Flyweight interface.
type ConcreteFlyweight struct {
	intrinsicState string
}

// NewConcreteFlyweight creates a new ConcreteFlyweight with the given intrinsic state.
func NewConcreteFlyweight(state string) *ConcreteFlyweight {
	return &ConcreteFlyweight{intrinsicState: state}
}

// Operation performs an operation using the extrinsic state and returns the result.
func (f *ConcreteFlyweight) Operation(extrinsicState string) string {
	return "IntrinsicState: " + f.intrinsicState + ", ExtrinsicState: " + extrinsicState
}
