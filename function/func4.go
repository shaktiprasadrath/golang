/*
Write a function with one variadic parameter
that finds the greatest number in a list of num-
bers.
*/
package main

import (
	"fmt"
	"sort"
)

//To check the greatest number
func checkGreatestNumber(x ...int) int {
	j := 0
	for _, v := range x {
		if j < v {
			j = v
		}
	}
	return j
}

func checkGreatestNumber_1(x ...int) int {
	sort.Ints(x)
	return x[len(x)-1]
}

//To check the mallest number
func checkSmallestNumber(x ...int) int {
	j := x[0]
	for _, v := range x {
		if j > v {
			j = v
		}
	}
	return j
}
func checkSmallestNumber_1(x ...int) int {
	sort.Ints(x)
	return x[0]
}

func main() {
	v := []int{56, 4, 44, 66, 77, 45, 78, 99, 234, 546}
	fmt.Println("The greatest number is: ", checkGreatestNumber(v...))
	fmt.Println("Another way to find greatest number is: ", checkGreatestNumber_1(v...))
	fmt.Println("The smallest number is: ", checkSmallestNumber(v...))
	fmt.Println("Another way to find smallest number is: ", checkSmallestNumber_1(v...))
	//checkSmallestNumber(v...)
}
