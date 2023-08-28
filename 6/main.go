package main

import (
	"fmt"
	"time"
)

var i int

func main() {
	// Создание канала для данных
	dataChan := make(chan string)
	// Создание канала для завершения
	done := make(chan struct{})

	// Запуск горутины, которая завершается при получении данных из канала done
	go func1(dataChan, done)
	// Заупск горутины, которая завершается при достижении определённого значения
	go func2()
	done <- struct{}{}
	time.Sleep(time.Second)
}
func func1(dataChan chan string, done chan struct{}) {
	for {
		// Проверяем канал на наличие сигнала остановки
		select {
		case data := <-dataChan:
			fmt.Println("Вычитывание данных из канала данных", data)
		case <-done:
			fmt.Println("Вычитывание из канала done")
			fmt.Println("Завершение горутины 1")
		}
	}
}

func func2() {
	for {
		i++
		if i == 10 {
			fmt.Println("Завершение горутины 2")
			return
		}
	}
}
