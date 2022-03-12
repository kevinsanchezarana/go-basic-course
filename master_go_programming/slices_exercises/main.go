package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
	exercise8()
}

func exercise1() {
	var cars = []string{"Audi", "Mercedes", "BMW"}
	for idx, car := range cars {
		fmt.Printf("Idx: %d and value %s\n", idx, car)
	}
}

func exercise2() {
	mySlice := []float64{1.2, 5.6}
	mySlice = append(mySlice, 6)
	a := float64(10)
	mySlice[0] = a
	mySlice = append(mySlice, 10.10)
	mySlice = append(mySlice, a)
	fmt.Println(mySlice)
}

func exercise3() {
	nums := []float64{1.5, 2.5, 3.5}
	nums = append(nums, 10.1)
	nums = append(nums, []float64{4.1, 5.5, 6.6}...) // or nums = append(nums, 4.1, 5.5, 6.6)
	fmt.Printf("Nums are %v\n", nums)
	n := []float64{7.5, 8.5}
	nums = append(nums, n...)
	fmt.Printf("Nums are %v\n", nums)
}

func exercise4() {
	if len(os.Args) < 2 || len(os.Args) > 10 {
		// log.Fatal("Number the parameters must be between 2 and 10") -> stop execution
		fmt.Println("Number the parameters must be between 2 and 10")
		return
	}

	var args = []string{}
	args = append(args, os.Args[1:]...)
	var product, sum float64
	for idx, arg := range args {
		var num, err = strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("Argument %s must be a float type", arg)
			continue
		}
		sum += num
		if idx == 0 {
			product = 1
		}
		product *= num
	}

	fmt.Printf("Expected output: Sum: %.0f, Product: %.0f\n", sum, product)
}

func exercise5() {
	sum := 0
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	for _, num := range nums[2 : len(nums)-2] {
		sum += num
		fmt.Printf("Number %d\n", num)
	}
	fmt.Printf("Total sum is %d\n", sum)
}

func exercise6() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	dst := make([]string, len(friends))
	copy(dst, friends) //generate a new slice without backing array
	dst[0] = "Kevin"
	fmt.Println(friends)
	fmt.Println(dst)
}

func exercise7() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	var dst = []string{}
	dst = append(dst, friends...) //generate a new slice without backing array
	dst[0] = "Kevin"
	fmt.Println(friends)
	fmt.Println(dst)
}

func exercise8() {
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	var newYears = []int{}
	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Printf("New Years %v\n", newYears)

}
