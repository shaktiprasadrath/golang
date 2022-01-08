package main

import (
	"fmt"
)

func main() {

	var b int = 10
	var p *int
	p = &b

	fmt.Println("Hello Shakti")
	fmt.Println("The value of b is: ", b)
	fmt.Println("The value of *p is: ", *p)

}
