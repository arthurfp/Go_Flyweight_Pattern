package flyweight

import "testing"

func TestConcreteFlyweight_Operation(t *testing.T) {
	flyweight := NewConcreteFlyweight("intrinsic")
	expected := "IntrinsicState: intrinsic, ExtrinsicState: extrinsic"
	result := flyweight.Operation("extrinsic")

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
