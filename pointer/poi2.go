package main

import "fmt"

type student struct {
	name   string
	age    int
	stream string
}

// copy value
func changeValue(val int) {
	p := val
	fmt.Println("func: changeValue ", p)
}

// Copy Address
func changeValueByAddress(p *int) {
	a := p
	fmt.Println("Func: changeValueByAddress ", a, *a)
}

//Copy value n return
func studentDetailsCopy() student {
	val := student{"Shakti", 21, "Physics"}
	fmt.Println(val)
	return val
}

//Copy address n return
func studentDetailsAdd() *student {
	val := student{"Shakti", 21, "Physics"}
	fmt.Println("Address of student ", &val)
	return &val
}

func main() {
	x := 10
	changeValue(x)
	y := 11
	changeValueByAddress(&y)

	//val := student{"Shakti", 21, "Physics"}
	fmt.Println(studentDetailsCopy())

	fmt.Println(studentDetailsAdd())

}
