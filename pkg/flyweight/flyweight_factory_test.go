package flyweight

import "testing"

func TestFlyweightFactory_GetFlyweight(t *testing.T) {
	factory := NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("state1")
	flyweight2 := factory.GetFlyweight("state2")
	flyweight3 := factory.GetFlyweight("state1")

	if flyweight1 == nil || flyweight2 == nil || flyweight3 == nil {
		t.Error("Expected non-nil flyweights")
	}

	if flyweight1 != flyweight3 {
		t.Error("Expected flyweight1 and flyweight3 to be the same instance")
	}

	if flyweight1 == flyweight2 {
		t.Error("Expected flyweight1 and flyweight2 to be different instances")
	}
}
