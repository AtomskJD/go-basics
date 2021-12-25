package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi my name is", p.Name)
}

type Android struct {
	Person
	Molel string
}

func main() {
	a := new(Android)
	a.Person.Name = "Jojo"
	a.Person.Talk()
}
