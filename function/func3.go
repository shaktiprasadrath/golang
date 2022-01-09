/*
 Write a function which takes an integer and
halves it and returns true if it was even or false
if it was odd. For example half(1) should return
(0, false) and half(2) should return (1,
true)
*/

package main

import "fmt"

func main() {
	half := func(x int) {
		if x%2 == 0 {
			fmt.Printf("%d is Even", x)
		} else {
			fmt.Printf("%d is Odd", x)
		}
	}

	half(10)

}
