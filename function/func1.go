package main

import "fmt"

func main() {
	x := []int{22, 44, 55, 66, 77}
	for p, q := range x {
		fmt.Println(p, q)
	}
}
