package main

import "fmt"

type Human struct {
	Age    int
	Name   string
	Gender string
	Weight float32
	Height float32
}

func (h Human) TellAboutMyself() {
	fmt.Printf("My Age=%v,Name=%v,Gender=%v,Weight=%v,Height=%v", h.Age, h.Name, h.Gender, h.Weight, h.Height)
}

type Action struct {
	Human //Встраивание структуры Human в Action
}

func main() {
	// Создание экземпляра структуры Action
	a := Action{Human{Age: 20, Name: "Eugene", Gender: "Male", Weight: 70, Height: 170}}
	// Вызов метода структуры Human из экземпляра структуры Action
	a.TellAboutMyself()
}
