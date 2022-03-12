package main

import "fmt"

func main() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {

	//Large via
	// var m map[string]int
	// m = map[string]int{}

	var m1 map[int]string
	fmt.Printf("%#v\n", m1)

	m2 := map[int]string{1: "One", 2: "Two"}
	m2[10] = "Abba"

	fmt.Printf("%#v\n", m2[10])
	fmt.Printf("%#v\n", m2[3])

}

func exercise2() {
	m1 := map[int]bool{}
	m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	s1 := fmt.Sprintf("%d", m2)
	s2 := fmt.Sprintf("%d", m3)

	fmt.Println(s1 == s2)
}

func exercise3() {
	m := map[int]bool{1: true, 2: false, 3: false}
	delete(m, 1)
	for k, v := range m {
		fmt.Printf("Key: %+v and Value: %+v\n", k, v)
	}
}
