package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	// Создание канала для квадратов чисел
	numSquares := make(chan int)

	// Конкурентное вычисление квадратов чисел
	for _, el := range nums {
		go calculateSquares(el, numSquares)
	}
	// Получение результата из канала и вывод
	for range nums {
		res := <-numSquares
		fmt.Println(res)
	}

}
func calculateSquares(num int, numSquares chan int) {
	numSquares <- num * num
}
