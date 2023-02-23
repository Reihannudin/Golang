package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHi(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	SayHi(eko)

	cat := Animal{
		Name: "Push",
	}
	SayHi(cat)
}
