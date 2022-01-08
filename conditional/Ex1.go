/*
 Write a program that prints out all the numbers
evenly divisible by 3 between 1 and 100. (3, 6, 9,
etc.)
*/

package main

import "fmt"

func main() {
	for p := 1; p <= 100; p++ {
		if p%3 == 0 {
			fmt.Println(p)
		}
	}

}
