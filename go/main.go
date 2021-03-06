package main

import (

	"example_simple"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	
	doSimple()
}

func doSimple()  {
	sm := example_simple.SimpleMessage{
		Id: 12345,
		IsSimple: true,
		Name: "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
	sm.Name = "I renamed you"
	fmt.Println(sm)

	fmt.Println("The ID is: ", sm.GetId())
}