package main

import (
	"fmt"
	"slices"
)

func kMin(arr []int, k int) []int {

	n := len(arr)
	k = min(k, n)

	arr_copy := make([]int, n)
	copy(arr_copy, arr)

	slices.Sort(arr_copy)

	array_k := make([]int, k)
	for i := 0; i < k; i++ {
		array_k[i] = arr_copy[i]
	}

	return array_k
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Реализовать функцию вычисления k  первых наименьших элементов числового массива.
	arr := []int{5, 3, 8, 4, 1, 2}

	result := kMin(arr, 3)

	fmt.Println(result) // [1 2 3]
}
