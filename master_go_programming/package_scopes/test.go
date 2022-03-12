package main

import "fmt"

func init() {
	fmt.Println("Init test.go file")
}

const testConstant = 100 //Its share in the package scope

func test() {
	fmt.Println("Package scope from another file")
}
