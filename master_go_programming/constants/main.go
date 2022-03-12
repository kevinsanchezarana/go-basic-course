package main

import "fmt"

//const are manage in complie time

func main() {
	const days, years int = 7, 1992
	fmt.Println(days, years)

	const (
		min1 = -500
		min2 = -300
		min3 = -100
	)

	fmt.Println(min1, min2, min3)

	const (
		max1 = -500
		max2
		max3
	)

	fmt.Println(max1, max2, max3)
}
