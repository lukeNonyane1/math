package math_module

import "golang.org/x/exp/constraints"

// package math_module provides a single utility to calculate the value to two numbers.

// Add is a function that takes in two number values and returns their sum.
// It follows the rules specified by https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}

type Number interface {
	constraints.Integer | constraints.Float
}
