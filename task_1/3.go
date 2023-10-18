package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Sin(x)
}

func buildArray(a, b float64, n int) []float64 {
	h := (b - a) / float64(n-1)
	result := make([]float64, n)
	for i := 0; i < n; i++ {
		x := a + float64(i)*h
		result[i] = f(x)
	}
	return result
}

func main() {
	/*
		Для заданного отрезка [a,b]  построить массив значений функции f(x).
		Реализовать [a,b]  построить массив значений функции f(x).
		Реализовать алгоритм  в виде отдельной функции.
	*/
	a, b := 0.0, math.Pi
	n := 10
	values := buildArray(a, b, n)
	fmt.Println("Массив значений:")
	for _, v := range values {
		fmt.Printf("%.5f ", v)
	}

	fmt.Println()
}
