package main

import "strconv"

var justString []rune

func someFunc() {
	v := []rune(createHugeString(1 << 10))
	// Необходима длина строки от 100 символов, чтоб взять первые 100 символов строки
	justString = v[:100]
	if len(v) >= 100 {
		justString = v[:100]
	}
}

func main() {
	someFunc()
}

func createHugeString(n int) string {
	return strconv.Itoa(n)
}
