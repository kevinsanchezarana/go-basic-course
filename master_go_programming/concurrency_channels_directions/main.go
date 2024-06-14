package main

import "fmt"

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	send(c1, "Hola")
	receiveAndSend(c1, c2)

	msg := <-c2
	fmt.Println(msg)

}

/*
*
@sender channel only write or send
@message string
*/
func send(c1 chan<- string, message string) {
	c1 <- message
}

/*
*
@sender channel only read or receive
@receiver channel only write or send
*/
func receiveAndSend(c1 <-chan string, c2 chan<- string) {
	msg := <-c1
	c2 <- msg
}
