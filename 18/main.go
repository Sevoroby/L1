package main

import (
	"fmt"
	"sync"
)

func main() {
	c := &concurrentCounter{}

	var wg sync.WaitGroup
	//Создаём 1000 горутин, в которых будет инкрементироваться структура-счётчик
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go testIncFunc(c, &wg)
	}

	wg.Wait()

	fmt.Println("Значение счетчика в конкурентной среде:", c.value)
}

type concurrentCounter struct {
	value int
	mu    sync.Mutex
}

func (c *concurrentCounter) inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
func testIncFunc(c *concurrentCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.inc()
}
