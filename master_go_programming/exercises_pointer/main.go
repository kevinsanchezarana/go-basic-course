package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	x := 10.10
	fmt.Printf("%v\n", x)
	ptr := &x
	fmt.Printf("%#v\n", ptr)  //Pointer to x
	fmt.Printf("%#v\n", *ptr) //Pointer value
	fmt.Printf("%#v\n", &ptr) //Pointer address itself
}

func exercise2() {
	x, y := 10, 2
	ptrx, ptry := &x, &y
	z := *ptrx / *ptry
	fmt.Printf("%#v\n", z)
}

func exercise3() {
	x, y := 5.5, 8.8
	fmt.Printf("BEFORE: %f, %f\n", x, y)
	swap(&x, &y)
	fmt.Printf("AFTER: %f, %f\n", x, y)
}

func swap(x, y *float64) {
	*x, *y = *y, *x
}
