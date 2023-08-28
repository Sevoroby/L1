package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{15, 20, 47, 2, 5, 13, 11}
	// Сортируем срез для двоичного поиска
	sort.Ints(nums)
	fmt.Println("Отсортированный массив:", nums)
	// Ищем элемент 13
	el := 15
	index, found := binarySearch(nums, el)
	if found {
		fmt.Printf("Найден элемент %d с индексом %d", el, index)
	} else {
		fmt.Println("Элемент не найден!")
	}
}

func binarySearch(nums []int, el int) (int, bool) {
	// Индекс левой границы
	left := 0
	// Индекс правой границы
	right := len(nums) - 1
	// Сдвигаем левую или правую границы, пока они не будут равны
	for left <= right {
		// Определяем центральный индекс
		mid := (left + right) / 2
		// Если центральный элемент равен искомому, то возвращаем результат, иначе сдвигаем одну из границ
		if nums[mid] == el {
			return mid, true
		} else if nums[mid] < el {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// Возвращаем false, если элемент не найден
	return 0, false
}
