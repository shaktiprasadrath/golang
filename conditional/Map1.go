//Example of Maps
package main

import "fmt"

func main() {
	x := map[int]string{1: "Ramu", 2: "Damu", 3: "Shamu"}
	fmt.Println(x)
	fmt.Println(x[1])
	if _, ok := x[2]; ok {
		fmt.Println(x[2])
	}
}
