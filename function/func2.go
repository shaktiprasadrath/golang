/*
sum is a function which takes a slice of numbers
and adds them together. What would its func-
tion signature look like in Go?
*/
package main

import "fmt"

func checkSlice(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	sl := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(checkSlice(sl))

}
