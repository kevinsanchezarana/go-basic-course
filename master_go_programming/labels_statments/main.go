package main

import "fmt"

func main() {
	outer := 19
	_ = outer

	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Michael"}
	friends := [2]string{"Mark", "Marry"}

	//label outer. Con be any name for label
outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("FOUND A FRIEND: %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break.")

	//the following piece of code creates a loop like a for statement does
	i := 0
loop: // label
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//ERROR:
	// 	goto todo
	//     x := 10
	//     fmt.Println(x)

	// todo:
	//     fmt.Println("something here")
}
