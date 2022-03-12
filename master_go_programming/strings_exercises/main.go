package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

func exercise1() {
	var name string = "Kevin"
	country := "Telde"
	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)
	fmt.Printf(`Your name: %s
	Country: %s
	`, name, country)
	fmt.Println("a) He says: \"Hello\"")
	fmt.Println(`C:\Users\a.txt`)
}

func exercise2() {
	var r rune = 'ă'
	fmt.Printf("%c\n", r)
	a, b := "ma", "m"
	result := a + b + string(r)
	fmt.Printf("%s\n", result)
}

func exercise3() {
	s1 := "țară means country in Romanian"

	// iterating over the string and print out byte by byte
	fmt.Printf("Bytes in string: ")
	for i := 0; i < len(s1); i++ {

		fmt.Printf("%v ", s1[i])
	}

	fmt.Println()

	// iterating over the string and print out rune by rune
	// and the byte index where the rune starts in the string
	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}

	fmt.Println()
}

func exercise4() {
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	// s1[0] = 'x' --> inmutable

	// printing the number of runes in the string
	fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))
}

func exercise5() {
	s := "你好 Go!"
	slice := []byte(s)
	fmt.Println(slice)
	for i, v := range slice {
		fmt.Printf("%d -> %d\n", i, v)
	}
}

func exercise6() {
	s := "你好 Go!"
	slice := []rune(s)
	fmt.Println(slice)
	for i, r := range slice {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}
}
