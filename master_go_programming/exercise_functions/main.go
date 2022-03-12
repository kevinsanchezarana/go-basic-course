package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%f\n", cube(3.))
	factorial, sum2 := f1(4)
	fmt.Printf("Factorial: %d, Sum: %d\n", factorial, sum2)
	fmt.Printf("myFunc: %d\n", myFunc("9"))
	fmt.Printf("sum: %d\n", sum(1, 2, 3, 4, 5, 6, 7))
	fmt.Printf("sum naked result: %d\n", sumNakedResult(1, 2, 3, 4, 5, 6, 7))
	animals := []string{"lion", "tiger", "bear"}
	fmt.Printf("searchItem: %v\n", searchItem(animals, "tiger"))
	fmt.Printf("searchItem: %v\n", searchItemInsensitive(animals, "Tiger"))
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
	number := printInt
	fmt.Printf("%T\n", number)
	number(8)
}

func cube(base float64) float64 {
	return math.Pow(base, 3)
}

func f1(n uint) (int, int) {
	factorial, sum := 1, 0
	for i := 1; i <= int(n); i++ {
		factorial *= i
		sum += i
	}
	return factorial, sum
}

func myFunc(n string) int {
	number1, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal("Argument must be an integer type")
	}
	twiceTimes := strings.Repeat(n, 2)
	treeTimes := strings.Repeat(n, 3)
	number2, _ := strconv.Atoi(twiceTimes)
	number3, _ := strconv.Atoi(treeTimes)
	return number1 + number2 + number3
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func sumNakedResult(numbers ...int) (total int) {
	for _, num := range numbers {
		total += num
	}
	return
}

func searchItem(haystack []string, needle string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

func searchItemInsensitive(haystack []string, needle string) bool {
	for _, v := range haystack {
		if strings.EqualFold(v, needle) {
			return true
		}
	}
	return false
}

func print(msg string) {
	fmt.Println(msg)
}

func printInt(number int) {
	fmt.Println(number)
}
