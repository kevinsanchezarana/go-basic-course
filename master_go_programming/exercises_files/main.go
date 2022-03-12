package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

func exercise1() {
	file, err := os.Create("exercises_files/info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func exercise2() {

	const originalPath string = "exercises_files/info.txt"
	const newPath string = "exercises_files/data.txt"

	// CHECKING IF FILE EXISTS
	_, err := os.Stat(originalPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}

	err = os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}

func exercise3() {
	const newPath string = "exercises_files/data.txt"
	err := os.Remove(newPath)
	if err != nil {
		log.Fatal(err)
	}
}

func exercise4() {
	data, err := ioutil.ReadFile("exercises_files/file4.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
}

func exercise5() {
	file, err := os.Open("exercises_files/file5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	iter := 0
	for scanner.Scan() {
		iter++
		fmt.Printf("Iter: %d, lines: %s\n", iter, scanner.Text())
	}

	// Checking for any possible errors:
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func exercise6() {
	file, err := os.Create("exercises_files/file6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("The Go gopher is an iconic mascot!")
	_, err = file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
}
