package main

import (
	"fmt"
	"math"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func sum(n1 float64, n2 float64, wg *sync.WaitGroup) {
	fmt.Printf("Sum is: %.2f\n", n1+n2)
	wg.Done()
}

func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	//Mutual exclusion
	m.Lock()
	*b += n
	m.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	//Mutual exclusion
	m.Lock()
	defer m.Unlock()
	*b -= n
	//m.Unlock() -> it is valid too
	wg.Done()
}

func main() {
	//Exercise 1
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello("Mr. Wick", &wg)
	wg.Wait()

	//Exercise 2
	wg.Add(3)
	for i := 0; i <= 2; i++ {
		go sum(3., 3., &wg)
	}
	wg.Wait()

	//Exercise 3
	wg.Add(1)
	go func(n float64, wg *sync.WaitGroup) {
		fmt.Printf("Square root number %0.f is: %6.f\n", n, math.Sqrt(n))
		wg.Done()
	}(4, &wg)
	wg.Wait()

	//Excersice 4
	wg.Add(50)
	for i := 100; i < 150; i++ {
		go func(n float64, wg *sync.WaitGroup) {
			fmt.Printf("Square root number %0.f is: %6.f\n", n, math.Sqrt(n))
			wg.Done()
		}(float64(i), &wg)
	}
	wg.Wait()

	//Excersice 5
	wg.Add(200)

	balance := 100 //is a shared resource and could have data race problems if several gouroutines write in it at once

	//If there are more than one simultaneous goroutine at the same time, at least one wait for another until the resource will be unlocked
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)

}
