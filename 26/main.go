package main

import (
	"fmt"
	"strings"
)

func isUniqueStr(str string) bool {
	// Мапа для хранения символов в нижнем регистре
	temp := make(map[string]struct{})

	for _, symbol := range str {
		// Приводим символ к нижнему регистру
		symbolLower := strings.ToLower(string(symbol))
		// Если в мапе есть данный символ, то возвращаем результат, что строка неуникальна, иначе записываем символ в мапу
		if _, inMap := temp[symbolLower]; inMap {
			return false
		} else {
			temp[symbolLower] = struct{}{}
		}
	}
	return true
}

func main() {
	// Строки для проверки
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	str4 := "тест"

	fmt.Println(str1, "-", isUniqueStr(str1))
	fmt.Println(str2, "-", isUniqueStr(str2))
	fmt.Println(str3, "-", isUniqueStr(str3))
	fmt.Println(str4, "-", isUniqueStr(str4))
}
