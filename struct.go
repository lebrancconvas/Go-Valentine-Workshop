package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {
	var Fern = Person{
		name: "Fern",
		age: 26,
	}
	fmt.Println(Fern)
}