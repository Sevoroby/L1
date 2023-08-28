package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Последовательность:", sequence)
	// Создаём мапу, где нужжны только ключи типа string, поскольку эти значения уникальны
	set := make(map[string]struct{})
	for _, el := range sequence {
		set[el] = struct{}{}
	}
	fmt.Println("Множество:")
	for el := range set {
		fmt.Print(el, " ")
	}
}
