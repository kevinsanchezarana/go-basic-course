package main

import "fmt"

func main() {
	language := "Golang"

	switch language {
	case "PHP":
		fmt.Println("PHP language :D")
	case "Java", "Python":
		fmt.Println("Java and Python languages :D")
	default:
		fmt.Println("Rest of languages :D")
	}

	n := 5
	// switch true {
	switch {
	case n%2 == 0:
		fmt.Println("Even number")
	case n%2 != 0:
		fmt.Println("Odd number")
	}
}
