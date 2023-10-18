package main

import (
	"fmt"

	"test.com/polynom"
)

func main() {
	/*
		Создать пакет polynom  предоставляющий интерфейс для символьных вычислений с полиномами
	*/

	p1 := polynom.New([]float64{2, 1})
	p2 := polynom.New([]float64{1, -3, 5})

	fmt.Println(p1) // 2x^1 + 1x^0
	fmt.Println(p2) // 5x^2 - 3x^1 + 1x^0

	y := p1.Calc(4)
	fmt.Println(y)

	p3 := p1.Add(p2)
	fmt.Println(p3) // 5x^2 + 2x^1 + 2x^0
}
