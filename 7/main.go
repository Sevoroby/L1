package main

import (
	"fmt"
	"sync"
)

func main() {
	N := 10000
	goroutineCount := 5
	m := make(map[int]int)
	mu := new(sync.RWMutex)

	var wg sync.WaitGroup
	// Добавление значений в вейт группу
	wg.Add(N * goroutineCount)
	// Запуск определённого количества горутин для записи в мапу
	for i := 0; i < goroutineCount; i++ {
		go testFunc(m, &wg, N, mu)
	}
	wg.Wait()
	// Вывод ключей и значений мапы
	for i, el := range m {
		fmt.Printf("m[%d]=%d\n", i, el)
	}
}
func testFunc(m map[int]int, wg *sync.WaitGroup, N int, mu *sync.RWMutex) {
	for i := 0; i < N; i++ {
		// Конкруентная запись в мапу
		writeToMap(m, i, mu)
		wg.Done()
	}
}
func writeToMap(m map[int]int, i int, mu *sync.RWMutex) {
	// Конкурентное считывание из мапы
	el := readFromMap(m, i, mu)
	mu.Lock()
	m[i] = el + i
	mu.Unlock()
}
func readFromMap(m map[int]int, i int, mu *sync.RWMutex) int {
	mu.RLock()
	el := m[i]
	mu.RUnlock()
	return el
}
