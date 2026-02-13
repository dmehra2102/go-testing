package basics

import "errors"

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Substract(a, b int) int {
	return a - b
}

func (c *Calculator) Multiply(a, b int) int {
	return a * b
}

func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func (c *Calculator) Power(base, exp int) int {
	if exp == 0 {
		return 1
	}

	result := base
	for i := 1; i < exp; i++ {
		result *= base
	}

	return result
}
