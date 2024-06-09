package flyweight

// FlyweightFactory creates and manages flyweight objects.
type FlyweightFactory struct {
	flyweights map[string]*ConcreteFlyweight
}

// NewFlyweightFactory creates a new FlyweightFactory.
func NewFlyweightFactory() *FlyweightFactory {
	return &FlyweightFactory{flyweights: make(map[string]*ConcreteFlyweight)}
}

// GetFlyweight returns a flyweight object for the given key, creating it if necessary.
func (f *FlyweightFactory) GetFlyweight(key string) *ConcreteFlyweight {
	if flyweight, exists := f.flyweights[key]; exists {
		return flyweight
	}
	flyweight := NewConcreteFlyweight(key)
	f.flyweights[key] = flyweight
	return flyweight
}
