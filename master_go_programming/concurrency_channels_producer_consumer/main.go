package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		x, open := <-c
		if !open {
			fmt.Println("channel closed", x)
			break
		}
		fmt.Println("value:", x)
	}
}
