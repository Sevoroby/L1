package main

import "fmt"

func main() {
	nums := []int{15, 20, 47, 2, 5, 13, 2}
	fmt.Println("Неотсортированный массив:", nums)
	fmt.Println("Отсортированный массив:", quickSort(nums))
}

func quickSort(nums []int) []int {
	// Если срез пустой или состоит из одного элемента, то возвращаем его
	if len(nums) <= 1 {
		return nums
	}
	// Определяем последний элемент среза как опорный, с которым будем сравнивать остальные элементы
	pivot := nums[len(nums)-1]
	// Срез для элементов, которые меньше опорного элемента
	left := make([]int, 0)
	// Срез для элементов, равных опорному
	equal := make([]int, 0)
	// Срез для элементов, которые больше опорного элемента
	right := make([]int, 0)

	// Проходимся по всем элементам среза и сравниваем с опорным, добавляя в соответствующие срезы
	for _, el := range nums {
		if el < pivot {
			left = append(left, el)
		} else if el > pivot {
			right = append(right, el)
		} else {
			equal = append(equal, el)
		}
	}
	// Рекурсивно выполняем функцию быстрой сортировки для срезов с меньшими и большими элементами
	left = quickSort(left)
	right = quickSort(right)

	// Объединяем отсортированный срез меньших элементов со срезом элементов, схожих с опорным
	left = append(left, equal...)
	// Затем объединяем с отсортированным срезом больших элементов
	res := append(left, right...)
	return res
}
