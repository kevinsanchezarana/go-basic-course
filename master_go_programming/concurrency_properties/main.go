/////////////////////////////////
// Getting Information
// Go Playground: https://play.golang.org/p/OHyIck2IwkS
/////////////////////////////////

package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func writeAFile() {
	file, err := os.OpenFile(
		"concurrency_properties/b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// WRITING BYTES TO FILE

	byteSlice := []byte("I learn Golang! 传 I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传I learn Golang! 传") // converting a string to a bytes slice
	bytesWritten, err := file.Write(byteSlice)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  // writing bytes to file.
	// It returns the no. of bytes written and an error value
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWritten)
}

func main() {
	// NumCPU returns the number of logical CPUs or cores usable by the current process.
	fmt.Println("No. of CPUs:", runtime.NumCPU()) // => No. of CPUs: 4

	// NumGoroutine returns the number of goroutines that currently exists.
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => No. of Goroutines: 1

	// GOOS and GOARCH are environment variables
	fmt.Println("OS:", runtime.GOOS)     // => OS: linux
	fmt.Println("Arch:", runtime.GOARCH) // => Arch: amd64

	//  GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns
	//  the previous setting.
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // => GOMAXPROCS: 4
	// If n < 1, it does not change the current setting.

	go writeAFile()

	fmt.Println("It is over")

}
