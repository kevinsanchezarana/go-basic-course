package main

import "fmt"

func main() {
	//Exercise 1
	var c1 chan float64
	//Unbuffered - SEND-ONLY
	c2 := make(chan<- rune)
	//Unbuffered - RECEIVE-ONLY
	c3 := make(<-chan rune)
	//Biderectional buffered channel
	c4 := make(chan int, 10)

	fmt.Printf("%+v\n", c1)
	fmt.Printf("%+v\n", c2)
	fmt.Printf("%+v\n", c3)
	fmt.Printf("%+v\n", c4)

	//Exercise 2
	c5 := make(chan string)
	go func() {
		c5 <- "Hi!!"
	}()

	fmt.Printf("%s\n", <-c5)

	//Exercise 3
	c := make(chan int)
	go func(n int) {
		c <- n
	}(100)

	fmt.Printf("%d\n", <-c)

	//Exercise 4
	for i := 1; i <= 50; i++ {
		go power(i, c)
	}

	for i := 1; i <= 50; i++ {
		fmt.Printf("Result Power is: %+v\n", <-c)
	}

	//Exercise 5
	for i := 1; i <= 50; i++ {
		go func(n int) {
			c <- n * n
		}(i)
	}

	for i := 1; i <= 50; i++ {
		fmt.Printf("Result Power is: %+v\n", <-c)
	}
}

func power(n int, c chan int) {
	c <- n * n
}
