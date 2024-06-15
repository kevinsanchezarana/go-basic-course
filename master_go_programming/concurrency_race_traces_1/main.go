package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	// Rate Tracing:  A data race occurs when two goroutines access the same variable concurrently and at least one of the accesses is a write.
	// We can check it adding the -race option:
	/*
		$ go test -race mypkg    // to test the package
		$ go run -race mysrc.go  // to run the source file
		$ go build -race mycmd   // to build the command
		$ go install -race mypkg // to install the package
	*/
	// raceUnprotectedGlobalVariableInWritingSharedVariableAtTheSameTime()
	// solutionRaceUnprotectedGlobalVariableInWritingSharedVariableAtTheSameTime()
	// raceOnLoopCounter()
	// solutionRaceOnLoopCounter()
	// data := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// raceSharedVariable(data)
	// solutionRaceSharedVariable(data)
}

func raceUnprotectedGlobalVariableInWritingSharedVariableAtTheSameTime() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func solutionRaceUnprotectedGlobalVariableInWritingSharedVariableAtTheSameTime() {
	var mutex sync.Mutex
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		mutex.Lock()
		m["1"] = "a" // First conflicting access.
		mutex.Unlock()
		c <- true
	}()
	mutex.Lock()
	m["2"] = "b" // Second conflicting access.
	mutex.Unlock()
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func raceOnLoopCounter() {
	var wg sync.WaitGroup
	wg.Add(20)
	// i is a shared variable
	for i := 0; i < 20; i++ { // we are writing in the main routine
		go func() {
			fmt.Println(i) // Not the 'i' you are looking for. We are reading the same value in another routine
			wg.Done()
		}()
	}
	wg.Wait()
}

func solutionRaceOnLoopCounter() {
	var wg sync.WaitGroup
	wg.Add(5)
	// i is a shared variable
	// Solution: using a copy instead of a shared variable
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(j)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func raceSharedVariable(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // The second conflicting write to err.
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res
}

func solutionRaceSharedVariable(data []byte) chan error {
	// Introduce new variables in the goroutines
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			var err1 error
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			_, err1 = f1.Write(data)
			res <- err1
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // The second conflicting write to err.
	if err != nil {
		res <- err
	} else {
		go func() {
			var err2 error
			_, err2 = f2.Write(data)
			res <- err2
			f2.Close()
		}()
	}
	return res
}
