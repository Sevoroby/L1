package main

import "fmt"

func main() {
	slice := []int{1, 10, 15, 7, 4, 15, 2, 8, 11, 14}
	fmt.Println("Слайс:", slice)
	i := 3
	// Создаём новый срез объединением двух срезов
	sliceRes := append(slice[:i], slice[i+1:]...)
	fmt.Printf("Слайс после удаления элемента с индексом %d: %v", i, sliceRes)
}
