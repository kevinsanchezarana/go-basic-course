package main

import (
	"fmt"
	"time"
)

func main() {
	fiftyFirstNumbersDivisibleBySeven1()
	fiftyFirstNumbersDivisibleBySeven2()
	fiftyFirstNumbersDivisibleBySeven3()
	fiveHundredNumbersDivisibleBySevenAndFive()
	everyYearsSinceMyBirthYear()
	changeToSwitchStatement()
}

func fiftyFirstNumbersDivisibleBySeven1() {
	fmt.Printf("fiftyFirstNumbersDivisibleBySeven1\n")
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Printf("%+v\n", i)
		}
	}
}

func fiftyFirstNumbersDivisibleBySeven2() {
	fmt.Printf("fiftyFirstNumbersDivisibleBySeven2\n")
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			if i < 1 || i > 50 {
				continue
			}
			fmt.Printf("%+v\n", i)
		}
	}
}

func fiftyFirstNumbersDivisibleBySeven3() {
	fmt.Printf("fiftyFirstNumbersDivisibleBySeven3\n")
	iterNumber := 0
	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			if iterNumber >= 3 {
				break
			}
			fmt.Printf("%+v\n", i)
			iterNumber++
		}
	}
}

func fiveHundredNumbersDivisibleBySevenAndFive() {
	fmt.Printf("fiveHundredNumbersDivisibleBySevenAndFive\n")
	for i := 1; i <= 500; i++ {
		if i%7 == 0 || i%5 == 0 {
			fmt.Printf("%+v\n", i)
		}
	}
}

func everyYearsSinceMyBirthYear() {
	fmt.Printf("everyYearsSinceMyBirthYear\n")
	myBirthYear := 1992
	currentYear := time.Now().Year()
	for myBirthYear <= currentYear {
		fmt.Println(myBirthYear)
		myBirthYear++
	}
}

func changeToSwitchStatement() {
	age := 29
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
