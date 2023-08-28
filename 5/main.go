package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Введите количество секунд")
	var N int
	fmt.Scan(&N)
	// Создание канала для данных
	dataChan := make(chan string)
	// Создание таймера, который в отдельном канале вернёт значение через заданное время
	timer := time.After(time.Duration(N) * time.Second)

	// Чтение данных в отдельной горутине
	go readData(timer, dataChan)
	// Отправка данных в канал данных с задержкой в 1 секунду
	for i := 0; i < N; i++ {
		dataChan <- "data"
		time.Sleep(time.Second * 1)
	}
}

func readData(timer <-chan time.Time, dataChan chan string) {
	for {
		select {
		case <-timer:
			fmt.Println("Завершение программы")
			return
		case <-dataChan:
			fmt.Println("Приняты данные из канала данных")
		}
	}
}
