package main

import (
	"fmt"
)

func f1() {
	fmt.Println("I am gassing")
}

func f2(x int) {
	fmt.Println("The value of int is: ", x)
}

func f3(a func()) {
	a()
}

func f4(a1 func()) func() {
	a1()
	return func() {
		fmt.Println("I am returning function")
	}
}

func main() {
	f1()
	p := f2
	p(10)

	q := func(s string) {
		fmt.Println("Hello ", s)
	}
	q("shakti")

	func() {
		fmt.Println("Hello GO")
	}()

	f3(f1)

	p1 := f4(f1)
	p1()

}
