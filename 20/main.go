package main

import (
	"fmt"
)

func main() {
	str := "snow dog sun"
	fmt.Println(reverseString(str))
}
func reverseString(str string) string {
	res := ""
	word := ""
	// Создание среза для слов
	words := []string{}
	// Преобразование в срез рун для итерации по символам строки
	strRunes := []rune(str)
	for i, symbol := range strRunes {
		// Если встречаем пробел, то добавляем слово в срез для слов
		if string(symbol) == " " {
			words = append(words, word)
			word = ""
		} else {
			word += string(symbol)
		}
		// Когда достигаем конца строки, то тоже добавляем слово в срез
		if len(strRunes)-1 == i {
			words = append(words, word)
		}
	}
	fmt.Println(words)
	// Выводим слова из среза в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i] + " "
	}

	return res
}
