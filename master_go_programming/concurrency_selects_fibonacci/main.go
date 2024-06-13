package main

import "fmt"

// Channels are tools to communicate info between several goroutines or the main function
func fibonacci(c, quit, iter chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // Writing or sending the x value to the channel
			x, y = y, x+y
		case i := <-iter: // Reading or receiving the i value to the channel
			fmt.Println("Iter: ", i)
		case <-quit: // Reading or receiving the channel
			fmt.Println("quit")
			return
			/*default:
			fmt.Println("default")
			// It's executed when there is no other case. It's nor blocking*/
		}

	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	iter := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // Reading or receiving the channel, and calling the select case
			iter <- i        // Writing or sending the i value into the channel
		}
		quit <- 0 // Writing or sending the 0 value into the channel
	}()
	fibonacci(c, quit, iter)
}
