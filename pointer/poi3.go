/*
Write a program that can swap two integers
( x := 1; y := 2; swap(&x, &y) should give you
x=2 and y=1 ).
*/
package main

import "fmt"

func swap(x *int, y *int) (int, int) {
	*x, *y = (*x + *y - *x), (*x + *y - *y)
	fmt.Println(*x, *y)
	return *x, *y
}
func main() {
	x := 2
	y := 5

	p, q := swap(&x, &y)
	fmt.Println(p, q)

}
