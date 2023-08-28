package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 10, 15, 20, 7, 4, 14, 2}
	// Создание канала для чисел
	numChan := make(chan int)
	// Создание канала для удвоенных чисел
	doubleNumChan := make(chan int)

	// Запуск первой горутины для записи чисел из массива в первый канал
	go sendToNumChan(nums, numChan)
	// Запуск второй горутины для получения чисел из первого канала и записи во второй
	go sendToDoubleNumChan(numChan, doubleNumChan)

	// Получение чисел из второго канала
	for el := range doubleNumChan {
		fmt.Printf("Получено число %d из канала doubleNumChan\n", el)
	}
}
func sendToNumChan(nums []int, numChan chan int) {
	for _, el := range nums {
		numChan <- el
	}
	close(numChan)
}
func sendToDoubleNumChan(numChan <-chan int, doubleNumChan chan int) {
	for el := range numChan {
		doubleNumChan <- el * 2
	}
	close(doubleNumChan)
}
