package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//** Use valid paths according to your OS. **//

	const path string = "files/a.txt"
	const path2 string = "files/b.txt"
	const path3 string = "files/aaa.txt"
	const path4 string = "files/aa.txt"
	// CREATING A FILE

	// os.Create() function creates a file if it doesn't already exist. If it exists, the file is truncated.
	// it returns a file descriptor which is a pointer to os.File and an error value.
	newFile, err := os.Create(path)

	// error handling
	if err != nil {
		// log the error and exit the program
		log.Fatal(err) // the idiomatic way to handle errors

	}

	// TRUNCATING A FILE
	err = os.Truncate(path, 0) //0 means completely empty the file.

	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// CLOSING THE FILE
	newFile.Close()

	// OPEN AND CLOSE AN EXISTING FILE
	file, err := os.Open(path) // open in read-only mode

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	//OPENING a FILE WITH MORE OPTIONS
	file, err = os.OpenFile(path, os.O_APPEND, 0644)
	// We can Use opening attributes individually or combined
	// using an OR between them
	// e.g. os.O_CREATE|os.O_APPEND
	// or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// GETTING FILE INFO
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat(path)

	p := fmt.Println
	p("File Name:", fileInfo.Name())        // => File Name: a.txt
	p("Size in bytes:", fileInfo.Size())    // => Size in bytes: 0
	p("Last modified:", fileInfo.ModTime()) // => Last modified: 2019-10-21 16:16:00.325037748 +0300 EEST
	p("Is Directory? ", fileInfo.IsDir())   // => Is Directory?  false
	p("Pemissions:", fileInfo.Mode())       // => Pemissions: -rw-r-----

	// CHECKING IF FILE EXISTS
	fileInfo, err = os.Stat(path2)
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}

	// RENAMING AND MOVING A FILE
	oldPath := path
	newPath := path3
	err = os.Rename(oldPath, newPath)
	// error handling
	if err != nil {
		log.Fatal(err)
	}

	// REMOVING A FILE
	err = os.Remove(path4)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
