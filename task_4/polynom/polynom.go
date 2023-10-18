package polynom

import (
	"fmt"
	"math"
)

// Полином
type Polynom struct {
	coeffs []float64
}

// Создание полинома из коэффициентов
func New(coeffs []float64) *Polynom {
	return &Polynom{coeffs}
}

// Вычисление полинома в точке x
func (p *Polynom) Calc(x float64) float64 {
	result := 0.0
	for i, c := range p.coeffs {
		result += c * math.Pow(x, float64(i))
	}
	return result
}

// Сложение полиномов
func (p *Polynom) Add(p2 *Polynom) *Polynom {

	size := max(len(p.coeffs), len(p2.coeffs))

	resCoeffs := make([]float64, size)

	for i := 0; i < size; i++ {
		if i < len(p.coeffs) {
			resCoeffs[i] += p.coeffs[i]
		}
		if i < len(p2.coeffs) {
			resCoeffs[i] += p2.coeffs[i]
		}
	}

	return &Polynom{resCoeffs}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Вывод полинома как строку
func (p *Polynom) String() string {
	str := ""
	for i, c := range p.coeffs {
		str += fmt.Sprintf("%gx^%d", c, i)
		if i < len(p.coeffs)-1 {
			str += " + "
		}
	}
	return str
}
