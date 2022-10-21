package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) newPerson(name string) *person {
	pa := person{name: name}
	pa.age = 42
	return &pa
}

func main() {
	earth := person{name: "earth", age: 12}

	pang := earth.newPerson("pang")

	fmt.Println("pang", pang)

}
