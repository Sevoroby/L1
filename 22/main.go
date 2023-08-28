package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаём числа в строковых представлениях
	str1 := "63456678907528967890"
	str2 := "54645432510432632210"

	a := new(big.Int)
	b := new(big.Int)

	// Преобразуем строку в число
	_, ok := a.SetString(str1, 10)
	if !ok {
		fmt.Println("Ошибка при преобразовании строки в число а")
		return
	}

	_, ok = b.SetString(str2, 10)
	if !ok {
		fmt.Println("Ошибка при преобразовании строки в число b")
		return
	}

	res := new(big.Int)
	res.Mul(a, b)
	fmt.Println("Результат при умножении:", res)

	res = new(big.Int)
	res.Div(a, b)
	fmt.Println("Результат при делении:", res)

	res = new(big.Int)
	res.Add(a, b)
	fmt.Println("Результат при сложении:", res)

	// Вычитание
	res = new(big.Int)
	res.Sub(a, b)
	fmt.Println("Результат при вычитании:", res)
}
