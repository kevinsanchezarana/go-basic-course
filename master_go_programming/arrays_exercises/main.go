package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	fmt.Println("Exercise1")
	var cities [2]string
	fmt.Println(cities)
	var grades = [3]float64{20.5, 30.6}
	fmt.Println(grades)
	var taskDone = [...]bool{true, false, true, false}
	fmt.Println(taskDone)
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	for _, v := range grades {
		fmt.Println(v)
	}
}

func exercise2() {
	fmt.Println("Exercise2")
	nums := [...]int{30, -1, -6, 90, -6}
	count := 0
	for _, num := range nums {
		if num > 0 && num%2 == 0 {
			count++
		}
	}
	fmt.Println(count)
}

func exercise3() {
	fmt.Println("Exercise3")
	myArray := [3]float64{1.2, 5.6}
	myArray[2] = 6
	a := float64(10)
	myArray[0] = a
	myArray[2] = 10.10
	fmt.Println(myArray)
}
