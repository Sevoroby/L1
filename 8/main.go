package main

import (
	"fmt"
	"math"
)

func main() {
	var setToValue bool
	var bit int8
	var num int64
	fmt.Println("Введите число int64")
	fmt.Scan(&num)
	fmt.Printf("Число %b в двоичном виде\n", num)
	fmt.Println("Введите бит, который нужно изменить")
	fmt.Scan(&bit)
	fmt.Println("Введите 0 или 1, чтобы установить указанный бит в это значение")
	fmt.Scan(&setToValue)
	res := changeBit(num, bit, setToValue)
	fmt.Printf("Число %b в двоичном виде после преобразования\n", res)
}

func changeBit(num int64, bit int8, setBitTo1 bool) int64 {
	if setBitTo1 {
		// Возводим двойку в степень нужного бита и выполняем пораздряную дизъюнкцию с получившимся числом
		return num | int64(math.Pow(2, float64(bit)))
	} else {
		// Возводим двойку в степень нужного бита и выполняем пораздряную конъюнкцию с инвертированным получившимся числом
		return num & (^int64(math.Pow(2, float64(bit))))
	}
}
