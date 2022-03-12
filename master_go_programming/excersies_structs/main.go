package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
}

func exercise1() {
	type Person struct {
		name    string
		age     int
		colours []string
	}

	me := Person{name: "Kevin", age: 29, colours: []string{"Blue", "Red"}}
	you := Person{name: "Judith", age: 26, colours: []string{"Violet", "Orange"}}
	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)
}

func exercise2() {
	type Person struct {
		name    string
		age     int
		colours []string
	}

	me := Person{name: "Kevin", age: 29, colours: []string{"Blue", "Red"}}
	me.name = "Andrei"
	you := Person{name: "Judith", age: 26, colours: []string{"Violet", "Orange"}}
	favoriteColours := you.colours
	fmt.Printf("%#v\n", favoriteColours)
	you.colours = append(you.colours, "Yellow")
	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)
}

func exercise3() {
	type Person struct {
		name    string
		age     int
		colours []string
	}

	me := Person{name: "Kevin", age: 29, colours: []string{"Blue", "Red"}}
	for _, colour := range me.colours {
		fmt.Printf("%s\n", colour)
	}
}

func exercise4() {
	type Grades struct {
		grade  float64
		course string
	}
	type Person struct {
		name    string
		age     int
		colours []string
		grades  Grades
	}
	me := Person{name: "Kevin", age: 29, colours: []string{"Blue", "Red"}, grades: Grades{grade: 7.5, course: "Golang"}}
	you := Person{name: "Judith", age: 26, colours: []string{"Violet", "Orange"}, grades: Grades{grade: 9.5, course: "PHP"}}
	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)
}
