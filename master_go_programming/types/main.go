package main

import "fmt"

type age int        //new type, int is the underlying type
type oldAge age     //new type, int (not age) is the underlying type
type veryOldAge age //new type, int (not age) is the underlying type

func main() {
	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	//var x uint
	//x = s1 // different types
	//x = unit(s1)

	var x uint = 10
	s3 := speed(x)
	fmt.Println(s3)

	//alias types
	var a uint8 = 10
	var b byte
	//byte i an alias of yubt8
	b = a
	_ = b

	type testAlias = float64
	var s5 testAlias = 4.5
	fmt.Printf("%+v\n", s5+2.5)

}
