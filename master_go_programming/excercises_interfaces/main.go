package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	//Excercise 1
	var v vehicle = car{licenceNo: "4543TYR", brand: "Peugeot"}
	fmt.Printf("License is: %#v - Name is: %#v\n", v.License(), v.Name())

	//Excercise 2
	var empty interface{} = 1
	fmt.Printf("%T\n", empty)
	empty = "Hola"
	fmt.Printf("%T\n", empty)
	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)
	empty = append(empty.([]int), 4) //assetion type -> access to implementation
	fmt.Println(empty)

	//Excercise 3
	var x interface{}
	x = cube{edge: 5}
	v2 := volume(x.(cube)) //assetion type -> access to implementation
	fmt.Printf("Cube Volume: %v\n", v2)
}
