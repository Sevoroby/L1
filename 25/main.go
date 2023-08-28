package main

import (
	"fmt"
	"time"
)

func sleep(dur time.Duration) {
	// На указанное количество времени создаётся ожидание
	<-time.After(dur)
}

func main() {
	fmt.Println("Начало программы")
	// Ожидание 3 секунды
	sleep(3 * time.Second)
	fmt.Println("Завершение программы")
}
