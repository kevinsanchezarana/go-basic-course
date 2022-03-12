package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%0.2f\n", m)
}
func exercise1(money money) {
	money.print()
}

func (m money) printStr() string {
	return fmt.Sprintf("%0.2f\n", m)
}
func exercise2(m money) string {
	return money.printStr(m)
}

type book struct {
	price float64
	title string
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func exercise3(b book) float64 {
	return b.vat()
}

func (b *book) changePrice(p float64) {
	(*b).price = p
}

func exercise4() {
	// book value
	bestBook := book{title: "The Trial by Kafka\n", price: 9.9}

	// changing the price
	//(&bestBook).changePrice(11.99)
	bestBook.changePrice(11.99) // or -> compiler do this job by us

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
}

func main() {
	var m money = 10.57657567
	exercise1(m)
	moneyStr := exercise2(m)
	fmt.Println(moneyStr)
	book := book{price: 100, title: "Hatty Potter"}
	vat := exercise3(book)
	fmt.Printf("%v\n", vat)
	exercise4()
}
