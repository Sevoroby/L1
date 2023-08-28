package main

import "fmt"

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(temperatures)
	// Создаём мапу с группами, где ключами являются целые числа, а значениями является срез вещественных чисел
	groups := make(map[int][]float32)
	// Проходимся по срезу с температурами
	for _, el := range temperatures {
		// Вычисляем группу
		group := (int)(el/10) * 10
		// Температуру в срез для данной группы
		groups[group] = append(groups[group], el)
	}
	// Выводим результат
	for i, el := range groups {
		fmt.Printf("%d:%.1f\n", i, el)
	}
}
