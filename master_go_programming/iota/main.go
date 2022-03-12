package main

import "fmt"

func main() {

	// iota is number generator for constants which starts from zero

	const (
		c1 = iota // 0
		c2 = iota // 1
		c3 = iota // 2
	)
	fmt.Println(c1, c2, c3)

	const (
		a1 = iota
		a2
		a3
	)
	fmt.Println(a1, a2, a3)

	const (
		b1 = iota * 2
		b2
		b3
	)
	fmt.Println(b1, b2, b3)

	const (
		x = -(iota + 2) // -2
		_               //skip value
		y               // -4
		z               // -5
	)
	fmt.Println(x, y, z)

}
