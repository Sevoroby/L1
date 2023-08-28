package main

import "fmt"

func main() {
	var a int = 10
	var b string = "dada"
	var c bool = true
	var d = make(chan string)
	determineType(a)
	determineType(b)
	determineType(c)
	determineType(d)
}

// Функция определения типа переменной
func determineType(testVar interface{}) {
	fmt.Printf("%T\n", testVar)
}
