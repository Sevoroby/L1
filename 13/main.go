package main

import "fmt"

func main() {
	a, b := 5, 17
	fmt.Printf("a=%d,b=%d\n", a, b)
	// Используем стандартный оператор множественного присваивания
	a, b = b, a
	fmt.Printf("a=%d,b=%d\n", a, b)
}
