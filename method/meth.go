package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

func (per person) print_person() {
	fmt.Println(per.name)
	fmt.Println(per.age)
	fmt.Println(per.gender)
}

func main() {
	p := person{
		name:   "Shakti",
		age:    36,
		gender: "Male"}

	p.print_person()
}
