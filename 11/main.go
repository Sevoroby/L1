package main

import (
	"fmt"
)

func getIntersection(set1 map[int]struct{}, set2 map[int]struct{}) []int {
	// Создание среза для результатов пересечения
	intersection := []int{}
	// Перебираем множество меньшей длины
	if len(set1) < len(set2) {
		for el := range set1 {
			// Если в другом множестве тоже есть такой элемент, то добавляем его в результат
			if _, inMap := set2[el]; inMap {
				intersection = append(intersection, el)
			}
		}
	} else {
		for el := range set2 {
			if _, inMap := set1[el]; inMap {
				intersection = append(intersection, el)
			}
		}
	}
	return intersection
}

func main() {
	// Создание и заполнение множества 1
	set1 := make(map[int]struct{})
	set1[3] = struct{}{}
	set1[2] = struct{}{}
	set1[1] = struct{}{}
	set1[5] = struct{}{}
	set1[4] = struct{}{}
	// Создание и заполнение множества 2
	set2 := make(map[int]struct{})
	set2[8] = struct{}{}
	set2[5] = struct{}{}
	set2[6] = struct{}{}
	set2[4] = struct{}{}
	set2[7] = struct{}{}
	set2[12] = struct{}{}
	set2[1] = struct{}{}

	fmt.Println("Множество 1:")
	for el := range set1 {
		fmt.Print(el, " ")
	}
	fmt.Println()
	fmt.Println("Множество 2:")
	for el := range set2 {
		fmt.Print(el, " ")
	}
	fmt.Println()

	intersection := getIntersection(set1, set2)

	fmt.Println("Пересечение:", intersection)

}
