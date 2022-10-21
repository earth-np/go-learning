package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// fmt.Println(person{"Bob", 20})

	// fmt.Println(person{name: "Alice", age: 30})

	// fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	man := newPerson(("Earth"))

	fmt.Println(man)

	man.age = 20

	fmt.Println(man)

	sp := man

	sp.age = 12

	fmt.Println(sp)

	s := person{name: "Sean", age: 50}

	fmt.Println(s)
	ee := &s

	ee.age = 999
	fmt.Println(s)
	fmt.Println(ee)
	// s := person{name: "Sean", age: 50}
	// fmt.Println(s.name)

	// sp := &s
	// fmt.Println(sp.age)

	// sp.age = 51
	// fmt.Println(sp.age)
}
