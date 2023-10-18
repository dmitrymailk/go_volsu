package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Последовательная версия
func dotSequential(a, b []float64) float64 {
	start := time.Now()
	// вычисление скалярного произведения
	var result float64

	for i := 0; i < len(a); i++ {
		result += a[i] * b[i]
	}

	duration := time.Since(start)
	fmt.Println("Последовательно:", duration)

	return result
}

func partialDot(a, b []float64) float64 {
	var result float64
	for i := range a {
		result += a[i] * b[i]
	}
	return result
}

// Параллельная версия
func dotParallel(a, b []float64) float64 {
	start := time.Now()
	// параллельное вычисление
	var goroutines = runtime.NumCPU()
	partLen := len(a) / goroutines

	var results = make([]float64, goroutines)

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			start := i * partLen
			end := start + partLen
			// вычисление для части
			results[i] = partialDot(a[start:end], b[start:end])
			wg.Done()
		}(i)
	}

	// Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
	wg.Wait()

	// суммирование результатов
	var result float64
	for _, res := range results {
		result += res
	}

	duration := time.Since(start)
	fmt.Println("Параллельно:", duration)

	return result
}

func randVector(n int) []float64 {
	vec := make([]float64, n)
	for i := 0; i < n; i++ {
		vec[i] = rand.Float64()
	}
	return vec
}

func main() {
	// 2. Реализовать параллельное вычисление скалярного произведения векторов.
	n := 100000000

	a := randVector(n)
	b := randVector(n)

	dotSequential(a, b)
	dotParallel(a, b)
}
