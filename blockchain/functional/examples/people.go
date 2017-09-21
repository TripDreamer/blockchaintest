package main

import (
	"fmt"
	f "github.com/izqui/functional"
)

type Person struct {
	Name   string
	Height int
}

func main() {

	people := []Person{Person{Name: "Mary", Height: 160}, Person{Name: "Isla", Height: 80}, Person{Name: "Sam"}}
	fmt.Println("People", people)

	knowingHeight := f.Filter(func(person Person) bool {
		return (person.Height > 0)
	}, people).([]Person)
	fmt.Println("Knowing height people", knowingHeight)

	average := f.Reduce(func(a float32, person Person) float32 {
		return a + (float32(person.Height) / float32(len(knowingHeight)))
	}, people, float32(0)).(float32)
	fmt.Println("Average height among people we know their height: ", average)
}
