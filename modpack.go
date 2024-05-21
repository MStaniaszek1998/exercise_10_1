package exercise101

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Add sums up two Integer or Float numbers and can return float or integer respectively
// It has two parameters: a - which is a first int or float and b - which is a second int or float
// Great website [MathIsFun]
// [MathIsFun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
