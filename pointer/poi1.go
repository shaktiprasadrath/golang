package main

import "fmt"

func main() {
	x := 7
	y := &x
	fmt.Println(x, y)
	fmt.Println(*y)
	*y = 10
	fmt.Println(x, y)
	fmt.Println(*y)

	var p *int = &x
	fmt.Println(p, *p, &p)

}
