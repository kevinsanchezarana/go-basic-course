package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func sum2(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0, 7, 2, 8, -9, 4, 0, 7, 2, 8, -9, 4, 0, 7, 2, 8, -9, 4, 0}

	start := time.Now()

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	/*res := sum2(s)
	fmt.Println(res)*/

	elapsed := time.Since(start)
	fmt.Printf("Sum took %s", elapsed)

}
