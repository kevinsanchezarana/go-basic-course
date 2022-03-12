package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%+v\n", os.Args)

	if len(os.Args) > 1 {
		var param1, err = strconv.ParseFloat(os.Args[1], 64)
		fmt.Println(param1)
		if err != nil {
			fmt.Println("The param1 is not a float64 valid!")
		}

		//Short if with asignaments. Must be decared before.
		//Initializate param1 for IF spoce and execute the second part to evaluate the conditions
		if param1, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
			fmt.Println(param1)
		} else {
			fmt.Println("The param1 is not a float64 valid!")
		}
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("The program must have one unique param")
	} else if km, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("The first param not is an integer, %v", err)
		_ = km
	}

}
