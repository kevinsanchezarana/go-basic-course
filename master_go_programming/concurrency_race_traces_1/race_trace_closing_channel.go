package main

func main() {
	// raceCase()
	solutionCase()
}

func raceCase() {
	c := make(chan struct{}) // or buffered channel

	// The race detector cannot derive the happens before relation
	// for the following send and close operations. These two operations
	// are unsynchronized and happen concurrently.
	go func() { c <- struct{}{} }()
	close(c) // We are closing the channel without waiting for the goroutine value
}

func solutionCase() {
	c := make(chan struct{}) // or buffered channel

	go func() { c <- struct{}{} }()
	<-c
	close(c)
}
