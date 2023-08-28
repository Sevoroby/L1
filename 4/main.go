package main

import (
	"fmt"
	"time"
)

var N int

func main() {
	fmt.Println("Введите количество воркеров")
	fmt.Scan(&N)

	// Создание канала для данных
	dataChan := make(chan int)

	// Создание и запуск воркеров в отдельных горутинах
	for i := 0; i < 1; i++ {
		go readData(i, dataChan)
	}

	// Постоянная запись значения счетчика в канал
	for i := 1; ; i++ {
		dataChan <- i
		time.Sleep(time.Millisecond * time.Duration(200))
		// При достижении значения (количество воркеров) * 10 останавливаем программу
		if i == N*10 {
			close(dataChan)
			break
		}
	}
	fmt.Println("Программа завершена")
}

func readData(id int, dataChan <-chan int) {
	for data := range dataChan {
		fmt.Printf("Чтеные данных %d в воркере %d \n", data, id)
	}
}
