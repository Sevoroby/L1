package main

import "fmt"

func main() {
	str := "главрыба"
	fmt.Println(reverse(str))
}
func reverse(str string) string {
	var res string
	// Преобразование в срез рун для итерации по символам строки
	strRune := []rune(str)
	for i := len(strRune) - 1; i >= 0; i-- {
		// Записываем символы в срез в обратном порядке
		res += string(strRune[i])
	}
	return res
}
