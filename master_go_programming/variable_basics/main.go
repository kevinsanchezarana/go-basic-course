package main

import "fmt"

//The short assigment only in the main scope x := 1
//varare manage in runtime

func main() {
	var age int = 30
	fmt.Println("Age:", age)

	var name = "Kevin"
	_ = name // this operator is used to mute the compiler errors. _ can be only on the left hand side of the = operator

	surname := "Sanchez"
	fmt.Println("Surname:", surname)

	car, cost := "Auidi", 50000
	fmt.Println(car, cost)

	car, year := "BMW", 2022
	fmt.Println(car, year) // at least should be new

	var opened bool = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int // swapping variables
	i, j = 5, 8
	i, j = j, i
	fmt.Println(i, j)

	sum := 5 + 4
	fmt.Println(sum)

}
