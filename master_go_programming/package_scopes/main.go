package main

import (
	"fmt" //imports are in file scope
	"master_go_programming/package_no_executable/numbers"
)

var arr [10]int

//Execute previous main method and you can declarate a lot.
func init() {
	fmt.Println("Init main.go file")
	arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func init() {
	fmt.Println("Init main2.go file")
}

//Execution: go run *.go
func main() {
	fmt.Println("Main execution :D")
	test()
	fmt.Println(testConstant)
	fmt.Println(numbers.IsEven(2))
	fmt.Printf("%+v\n", arr)
}
