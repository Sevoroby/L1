package main

import (
	"fmt"
)

// Интерфейс для способности говорить
type Talker interface {
	Talk() float64
}

// Структура человека
type Human struct {
	Name   string
	Age    uint8
	Gender string
}

func (h *Human) Talk() {
	fmt.Printf("%s говорит...\n", h.Name)
}

// Структура животного
type Animal struct {
	Name  string
	Age   uint8
	Breed string
}

// Адаптер для животного
type AnimalAdapter struct {
	Animal
}

// Метод адаптера, который добавляет животному возможность говорить
func (aa *AnimalAdapter) Talk() {
	fmt.Printf("%s по имени %s говорит...\n", aa.Breed, aa.Name)
}

func main() {
	human1 := &Human{
		Name:   "Андрей",
		Age:    25,
		Gender: "Male",
	}

	human1.Talk()

	// Создание объекта структуры адаптера
	animal1 := &AnimalAdapter{Animal{
		Name:  "Бобик",
		Age:   5,
		Breed: "Доберман",
	},
	}
	// С помощью структуры адаптера мы можем вызвать метод Talk для животного
	animal1.Talk()
}
