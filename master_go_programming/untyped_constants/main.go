package main

import "fmt"

func main() {

	const a int = 10
	const b int = 2
	const c int = a * b

	fmt.Println(c)

	const x = 5 // untype const

	var i int = x     // similar to int(x)
	var j float64 = x // similar to float64(x)
	var p byte = x    // similar to byte(x)

	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", j)
	fmt.Printf("%T\n", p)

}
