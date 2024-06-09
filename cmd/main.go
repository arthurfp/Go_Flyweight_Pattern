package main

import (
	"flyweight-pattern-go/pkg/flyweight"
	"fmt"
)

func main() {
	factory := flyweight.NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("state1")
	flyweight2 := factory.GetFlyweight("state2")
	flyweight3 := factory.GetFlyweight("state1")

	fmt.Println(flyweight1.Operation("extrinsic1"))
	fmt.Println(flyweight2.Operation("extrinsic2"))
	fmt.Println(flyweight3.Operation("extrinsic3"))

	// Verify that flyweight1 and flyweight3 are the same instance
	fmt.Printf("flyweight1 and flyweight3 are the same instance: %v\n", flyweight1 == flyweight3)
}
